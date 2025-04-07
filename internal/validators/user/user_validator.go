package user

import (
	"fmt"

	"github.com/josevitorrodriguess/client-manager/internal/utils"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"role"`
}

func (ur *UserRequest) IsValid() (bool, error) {
	if !(utils.MinChars(ur.Name, 5) && utils.MaxChars(ur.Name, 100)) {
		return false, fmt.Errorf("name must have between 5 and 100 characters")
	}
	if !utils.Matches(ur.Email, utils.EmailRegex) {
		return false, fmt.Errorf("invalid email")
	}
	if !(utils.MinChars(ur.Password, 8) && utils.MaxChars(ur.Password, 100)) {
		return false, fmt.Errorf("password must have between 8 and 100 characters")
	}

	return true, nil
}

type UserRequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (url *UserRequestLogin) IsValid() (bool, error) {
	if !utils.Matches(url.Email, utils.EmailRegex) {
		return false, fmt.Errorf("invalid email")
	}
	if !(utils.MinChars(url.Password, 8) && utils.MaxChars(url.Password, 100)) {
		return false, fmt.Errorf("password must have between 8 and 100 characters")
	}
	return true, nil
}

type UserResponse struct{}
