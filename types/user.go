package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost   = 12
	minFristName = 2
	minLastName  = 2
	minPassword  = 6
)

func (params CreateUserParams) Validate() []string {
	errors := []string{}
	if len(params.FirstName) < minFristName {
		errors = append(errors, fmt.Sprintf("firstName length should be at least %d characters", minFristName))
	}
	if len(params.LastName) < minLastName {
		errors = append(errors, fmt.Sprintf("lastName length should be at least %d characters", minLastName))
	}
	if len(params.Password) < minPassword {
		errors = append(errors, fmt.Sprintf("password length should be at least %d characters", minPassword))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Sprintf("email in invalid"))
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // omitempty for dont show id if is empty
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
