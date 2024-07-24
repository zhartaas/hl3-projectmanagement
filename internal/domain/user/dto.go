package user

import (
	"errors"
	"net/http"
)

type Request struct {
	Name  string `json:"name" example:"zhartas"`
	Email string `json:"email" example:"zhartas@gmail.com"`
	Role  string `json:"role" example:"administrator"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Name == "" {
		return errors.New("name: cannot be blank")
	}

	if s.Email == "" {
		return errors.New("email:cannot be blank")
	}

	if s.Role != "administrator" && s.Role != "manager" && s.Role != "developer" {
		return errors.New("role: invalid role")
	}

	return nil
}

type Response struct {
}
