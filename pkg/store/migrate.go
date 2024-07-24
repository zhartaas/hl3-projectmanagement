package store

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(dsn string) (err error) {
	migrations, err := migrate.New("file://migrations/postgre", dsn)
	if err != nil {
		return
	}
	if err = migrations.Down(); err != nil {
		return
	}
	if err = migrations.Up(); err != nil && err != migrate.ErrNoChange {
		return
	}
	fmt.Println("Migration: ", err)
	return nil
}
