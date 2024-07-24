package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hl3-projectmanagement/internal/domain/project"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(project project.Entity) (err error) {
	query := `INSERT INTO 
	projects (id, title, description, start_date, end_date, manager_id) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	args := []any{project.ID, project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *ProjectRepository) GetProjects() (projects []project.Entity, err error) {
	query := `SELECT * FROM projects`

	rows, err := r.db.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var project project.Entity
		err = rows.StructScan(&project)
		if err != nil {
			return
		}

		projects = append(projects, project)
	}

	return
}

func (r *ProjectRepository) GetProject(id string) (project project.Entity, err error) {
	query := `SELECT * FROM projects WHERE id = $1`

	row := r.db.QueryRowx(query, id)

	row.StructScan(&project)

	return
}

func (r *ProjectRepository) UpdateProject(id string, project project.Entity) (err error) {
	query := `UPDATE projects 
	SET title = $1, description = $2, start_date = $3, end_date = $4, manager_id = $5
	WHERE id = $6`

	args := []any{project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID, id}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *ProjectRepository) DeleteProject(id string) (err error) {
	query := `DELETE FROM projects WHERE id = $1`

	_, err = r.db.Exec(query, id)

	return
}

func (r *ProjectRepository) SearchProject(params map[string]string) (tasks []project.Entity, err error) {
	query := `SELECT * FROM projects WHERE 1=1`

	args := make([]any, 0, len(params))
	for key, value := range params {
		if value != "" {
			query += fmt.Sprintf(" AND %s = $%v ", key, len(args)+1)
			args = append(args, value)
		}
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task project.Entity
		err = rows.StructScan(&task)
		if err != nil {
			return
		}

		tasks = append(tasks, task)
	}

	return
}
