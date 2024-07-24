package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hl3-projectmanagement/internal/domain/task"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetUserTasks(id string) (tasks []task.Entity, err error) {
	query := `SELECT * FROM tasks WHERE responsible_id = $1`
	args := []any{id}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task task.Entity
		err = rows.StructScan(&task)
		if err != nil {
			return
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	return
}

func (r *TaskRepository) CreateTask(req task.Entity) (err error) {
	query := `INSERT INTO 
    tasks (id, title, description, priority, status, responsible_id, project_id, completion_date) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	args := []any{req.ID, req.Title, req.Description, req.Priority, req.Status, req.ResponsibleID, req.ProjectID, req.CompletionDate}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *TaskRepository) ListTasks() (tasks []task.Entity, err error) {
	query := `SELECT * FROM tasks`

	rows, err := r.db.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var task task.Entity
		err = rows.StructScan(&task)
		if err != nil {
			return
		}

		tasks = append(tasks, task)
	}
	err = rows.Err()

	return
}

func (r *TaskRepository) GetTask(id string) (task task.Entity, err error) {
	query := `SELECT * FROM tasks WHERE id = $1`
	args := []any{id}

	row := r.db.QueryRowx(query, args...)

	row.StructScan(&task)

	return
}

func (r *TaskRepository) UpdateTask(id string, req task.Entity) (err error) {
	query := `UPDATE tasks SET title = $1, description = $2, priority = $3, status = $4, responsible_id = $5, project_id = $6, completion_date = $7 WHERE id = $8`
	args := []any{req.Title, req.Description, req.Priority, req.Status, req.ResponsibleID, req.ProjectID, req.CompletionDate, id}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *TaskRepository) DeleteTask(id string) (err error) {
	query := `DELETE FROM tasks WHERE id = $1`
	args := []any{id}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *TaskRepository) Search(params map[string]string) (tasks []task.Entity, err error) {
	query :=
		`SELECT * FROM tasks
		 WHERE 1=1
`

	args := make([]any, 0, len(params))
	for key, value := range params {
		if value != "" {
			query += fmt.Sprintf(" AND %s = $%v ", key, len(args)+1)
			args = append(args, value)
		}
	}

	rows, err := r.db.Queryx(query, args...)

	for rows.Next() {
		var task task.Entity
		err = rows.StructScan(&task)
		if err != nil {
			return
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	return
}
