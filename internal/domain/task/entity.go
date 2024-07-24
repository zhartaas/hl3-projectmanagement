package task

import "time"

type Entity struct {
	ID             string    `json:"id" db:"id"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	Priority       string    `json:"priority" db:"priority"`
	Status         string    `json:"status" db:"status"`
	ResponsibleID  string    `json:"responsible_id" db:"responsible_id"`
	ProjectID      string    `json:"project_id" db:"project_id"`
	CreationDate   time.Time `json:"created_at" db:"creation_date"`
	CompletionDate time.Time `json:"completed_at,omitempty" db:"completion_date"`
}
