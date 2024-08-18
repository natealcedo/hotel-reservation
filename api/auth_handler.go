package api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

const invalidCredentialsMsg = "invalid credentials"

type AuthParams struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(store db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: store,
	}
}

type genericResponse struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func invalidCredentials(c *fiber.Ctx) error {
	return c.Status(401).JSON(
		genericResponse{
			Type: "error",
			Msg:  invalidCredentialsMsg,
		})
}

// A handler should ony do: -> Comparable to a MVC pattern. The handler is the controller. We're prolly gonna
// need a business service layer
// - serialization/deserialization from request/response
// - do some data fetching
// - call some business logic
// - return the data back to the client
func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	// When using body parser, the variable has to be a struct, not a pointer to a struct
	// When passing the variable to c.BodyParser, it has to be a pointer to the variable
	// Not doing so will result in a nil dereference error
	var authParams AuthParams
	if err := c.BodyParser(&authParams); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(), authParams.Email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return invalidCredentials(c)
		}
		return err
	}

	if !types.IsValidPassword(user.EncryptedPassword, authParams.Password) {
		return invalidCredentials(c)
	}

	resp := &AuthResponse{
		User:  user,
		Token: CreateTokenFromUser(user),
	}

	return c.JSON(resp)
}

func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4)
	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("error signing token", err)
	}
	return tokenStr
}
