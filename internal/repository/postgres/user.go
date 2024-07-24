package postgres

import (
	"github.com/jmoiron/sqlx"
	"hl3-projectmanagement/internal/domain/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(data user.Entity) (err error) {
	query := `
	INSERT INTO users (id, name, email, role) 
	VALUES ($1,$2,$3,$4)  
`
	args := []any{data.ID, data.Name, data.Email, data.Role}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return
	}

	return nil
}

func (r *UserRepository) GetAll() (dest []user.Entity, err error) {
	query := `
	SELECT * FROM users
`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		user := user.Entity{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.DateOfRegister, &user.Role)
		if err != nil {
			return
		}
		dest = append(dest, user)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (r *UserRepository) GetByID(id string) (user user.Entity, err error) {
	query := `
	SELECT * FROM users
	WHERE id=$1
`
	args := []any{id}

	row := r.db.QueryRowx(query, args...)
	row.StructScan(&user)

	return
}

func (r *UserRepository) Update(id string, req user.Request) (err error) {
	query := `
	UPDATE users
	SET name=$1, email=$2, role=$3
	WHERE id = $4
`

	args := []any{req.Name, req.Email, req.Role, id}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *UserRepository) Delete(id string) (err error) {
	query := `
	DELETE FROM users 
	WHERE ID = $1
`
	args := []any{id}

	_, err = r.db.Exec(query, args...)

	return
}

func (r *UserRepository) Search(param string, searchByName bool) (user user.Entity, err error) {
	query := `
	SELECT * FROM users
`
	if searchByName {
		query += `WHERE name = $1`
	} else {
		query += `WHERE email = $1`
	}

	args := []any{param}

	row := r.db.QueryRowx(query, args...)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.DateOfRegister, &user.Role)

	return
}
