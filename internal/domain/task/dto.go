package task

import (
	"errors"
	"net/http"
	"time"
)

type Request struct {
	Title         string `json:"title" example:"add new feature" extensions:"x-order=1"`
	Description   string `json:"description" example:"add new feature to the project" extensions:"x-order=2"`
	Priority      string `json:"priority" example:"high" extensions:"x-order=3"`
	Status        string `json:"status" example:"in_progress" extensions:"x-order=4"`
	ResponsibleID string `json:"responsible_id" example:"id" extensions:"x-order=5"`
	ProjectID     string `json:"project_id" example:"id" extensions:"x-order=6"`
	CompletionAt  string `json:"completion_at" example:"2006-01-02" extensions:"x-order=7"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Title == "" {
		return errors.New("name: cannot be blank")
	}

	if s.Description == "" {
		return errors.New("description: cannot be blank")
	}

	if s.Priority != "high" && s.Priority != "medium" && s.Priority != "low" {
		return errors.New("priority: invalid priority")
	}

	if s.Status != "new" && s.Status != "in_progress" && s.Status != "completed" {
		return errors.New("status: invalid status")
	}

	// check if completion_at is valid
	if _, err := time.Parse(time.DateOnly, s.CompletionAt); err != nil {
		return errors.New("completion_at: invalid date")
	}

	return nil
}
