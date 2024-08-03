package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func JWTAuthentication(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok || len(token) == 0 {
		return fmt.Errorf("unauthorized")
	}

	if err := parseToken(token[0]); err != nil {
		fmt.Println("error parsing token", err)
		return fmt.Errorf("unauthorized")
	}

	if err := c.Next(); err != nil {
		return err
	}

	return nil

}

func parseToken(tokenStr string) error {
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
		return unauthorized
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

	return unauthorized
}
