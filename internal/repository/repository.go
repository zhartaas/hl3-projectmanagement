package repository

import (
	"hl3-projectmanagement/internal/repository/postgres"
	"hl3-projectmanagement/pkg/store"
)

// this file isn't neccessary in my project because I am using only postgresql
// if I would have different db, I can use this file to determine which db I am using.
// using interfaces it is convenient to send as a parameter interface instead of structure,
// and this structure implements interface in different db

type Repository struct {
	postgres store.SQLX

	User    postgres.UserRepository
	Task    postgres.TaskRepository
	Project postgres.ProjectRepository
}
