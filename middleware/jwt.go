package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/natealcedo/hotel-reservation/db"
	"os"
	"time"
)

func JWTAuthentication(store db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok || len(token) == 0 {
			fmt.Println("token not found")
			return fmt.Errorf("unauthorized")
		}

		claims, err := validateToken(token[0])
		if err != nil {
			fmt.Println("error parsing token", err)
			return fmt.Errorf("unauthorized")
		}

		// Check token expiration
		// Refactor this to have proper time expiration handling
		expiresStr, ok := claims["expires"].(string)
		if !ok {
			return fmt.Errorf("unauthorized")
		}
		expires, err := time.Parse(time.RFC3339, expiresStr)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}
		if time.Now().After(expires) {
			return fmt.Errorf("token expired")
		}

		userID := claims["id"].(string)
		user, err := store.GetUserById(c.Context(), userID)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}

		// Set authenticated user in context value
		c.Context().SetUserValue("user", user)

		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	unauthorized := fmt.Errorf("unauthorized")
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("unexpected signing method: %v\n", token.Header["alg"])
			return nil, unauthorized
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, unauthorized
	}

	if !parsedToken.Valid {
		fmt.Println("invalid token")
		return nil, unauthorized
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("invalid claims")
		return nil, unauthorized
	}

	return claims, nil
}
