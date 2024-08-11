package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName         string             `json:"firstName" bson:"firstName"`
	LastName          string             `json:"lastName" bson:"lastName"`
	Email             string             `json:"email" bson:"email"`
	EncryptedPassword string             `json:"-" bson:"encryptedPassword"`
	IsAdmin           bool               `json:"isAdmin"`
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
}

func isValidEmail(email string) bool {
	emailRegexPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	return emailRegexPattern.MatchString(email)
}

func (params UpdateUserParams) Validate() map[string]string {
	errors := make(map[string]string)
	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
	}

	return errors
}

func (params CreateUserParams) Validate() map[string]string {
	errors := make(map[string]string)
	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}

	if !isValidEmail(params.Email) {
		errors["email"] = fmt.Sprintf("email address %s is not valid", params.Email)
	}

	return errors
}

func NewUserFromParams(params *CreateUserParams) (*User, error) {
	encryptedPw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptedPw),
	}, nil
}

func IsValidPassword(password, encryptedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(encryptedPassword)) == nil
}
