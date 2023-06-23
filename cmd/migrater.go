package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "money-keeper"
	password = "123456"
	dbname   = "money-keeper-dev"
)

func CreateMigrate() {

}

func MigrateUp() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		log.Panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/internal/migrations", dir),
		"postgres",
		driver)
	if err != nil {
		fmt.Printf("⚡️ migrations init fail <%s>", err)
	}
	err = m.Up()
	fmt.Println("✨ Success migrate")
	if err != nil {
		fmt.Printf("⚡️ migrations up fail <%s>", err)
	}
}
