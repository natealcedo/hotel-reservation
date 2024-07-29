package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName         string             `json:"firstName" bson:"firstName"`
	LastName          string             `json:"lastName" bson:"lastName"`
	Email             string             `json:"email" bson:"email"`
	EncryptedPassword string             `json:"-" bson:"encryptedPassword"`
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
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
