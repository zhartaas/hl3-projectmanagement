package store

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLX struct {
	Client *sqlx.DB
}

func New(dsn string, insertExampleValues bool) (store SQLX, err error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}

	if err = Migrate(dsn); err != nil {
		err = errors.New("error in migrate:" + err.Error())
	}
	store.Client = db
	fmt.Println("Migrate successfully")

	if insertExampleValues {
		err = store.InsertExampleValues()
		if err != nil {
			errors.Join(errors.New("Insert example values error: "), err)
		}
	}

	return
}

func (s *SQLX) Close() {
	s.Client.Close()
}

func (s *SQLX) InsertExampleValues() (err error) {

	insertUserQuery := `
	INSERT INTO users (id, name, email, role)
	VALUES ('11000000-0000-0000-0000-000000000000', 'example user', 'example.user@gmail.com', 'manager')
`
	insertProjectQuery := `
	INSERT INTO projects (id, title, description, start_date, end_date, manager_id)
	VALUES ('22000000-0000-0000-0000-000000000000', 'example project', 'example description', '2021-01-01', '2021-12-31', '11000000-0000-0000-0000-000000000000')
`
	insertTaskQuery := `
	INSERT INTO tasks (id, title, description, priority, status, responsible_id, project_id, creation_date, completion_date)
	VALUES ('33000000-0000-0000-0000-000000000000', 'example task', 'example description', 'low', 'in_progress', '11000000-0000-0000-0000-000000000000', '22000000-0000-0000-0000-000000000000', '2021-01-01', '2021-12-31')
`
	_, err = s.Client.Exec(insertUserQuery)
	if err != nil {
		return
	}
	_, err = s.Client.Exec(insertProjectQuery)
	if err != nil {
		return
	}

	_, err = s.Client.Exec(insertTaskQuery)
	if err != nil {
		return
	}

	return
}
