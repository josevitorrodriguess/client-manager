package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Auth struct {
	Id       *uuid.UUID `json:"id"`
	Email    string     `json:"email" validate:"required,email"`
	Password string     `json:"password" validate:"required"`
}

func ValidateAuth(auth *Auth) error {
	validate := validator.New()
	err := validate.Struct(auth)
	if err != nil {
		return err
	}
	return nil
}
