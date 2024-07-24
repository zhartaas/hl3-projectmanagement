package project

import (
	"errors"
	"net/http"
	"time"
)

type Request struct {
	Title       string `json:"title" example:"ai startup" extensions:"x-order=1"`
	Description string `json:"description" example:"text generator" extensions:"x-order=2"`
	StartedAt   string `json:"started_at" example:"2006-01-02" extensions:"x-order=3"`
	EndAt       string `json:"end_at" example:"2006-01-02" extensions:"x-order=4"`
	ManagerID   string `json:"manager_id" example:"id" extensions:"x-order=5"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Title == "" {
		return errors.New("title: cannot be blank")
	}

	if s.Description == "" {
		return errors.New("description: cannot be blank")
	}

	if _, err := time.Parse(time.DateOnly, s.StartedAt); err != nil {
		return errors.New("started_at: invalid date")
	}

	if _, err := time.Parse(time.DateOnly, s.EndAt); err != nil {
		return errors.New("end_at: invalid date")
	}

	return nil
}
