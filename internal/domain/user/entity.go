package user

import "time"

type Entity struct {
	ID             string
	Name           string
	Email          string
	DateOfRegister time.Time
	Role           string // роль пользователя в системе (например, администратор, менеджер, разработчик).
}
