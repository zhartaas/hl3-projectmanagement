package app

import (
	"fmt"
	"hl3-projectmanagement/internal/handler"
	"hl3-projectmanagement/internal/service/management"
	"hl3-projectmanagement/pkg/server"
	"hl3-projectmanagement/pkg/store"
	"log"

	"github.com/joho/godotenv"
	"os"
)

// /internal/app: this section may include any initialization code that needs to be executed before the application starts.
// For example, setting up configuration, connecting to databases, or initializing logging.

func Run() {
	log.SetFlags(log.Llongfile)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	postgreUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := "postgres"
	port := "5432"

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, postgreUser, password, dbName)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgreUser, password, host, port, dbName)

	insertExampleValues := true
	db, err := store.New(dsn, insertExampleValues)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	if err = db.Client.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	managementService := management.New(db.Client)

	handlers, err := handler.New(managementService)
	if err != nil {
		log.Fatal(err)
		return
	}

	serv, err := server.New(handlers.HTTP, "8080")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = serv.Run(); err != nil {
		log.Fatal(err)
		return
	}
	quit := make(chan os.Signal, 1) //

	<-quit
}
