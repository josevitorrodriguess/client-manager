package auth

import "github.com/go-playground/validator/v10"

type Auth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ValidateAuth(auth *Auth) error {
	validate := validator.New()
	err := validate.Struct(auth)
	if err != nil {
		return err
	}
	return nil
}
