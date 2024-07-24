package management

import (
	"github.com/jmoiron/sqlx"
	"hl3-projectmanagement/internal/repository/postgres"
)

type Service struct {
	UserRepository    *postgres.UserRepository
	TaskRepository    *postgres.TaskRepository
	ProjectRepository *postgres.ProjectRepository
}

func New(db *sqlx.DB) (s *Service) {
	s = &Service{
		UserRepository:    postgres.NewUserRepository(db),
		TaskRepository:    postgres.NewTaskRepository(db),
		ProjectRepository: postgres.NewProjectRepository(db),
	}
	return
}
