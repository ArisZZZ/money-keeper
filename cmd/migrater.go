package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

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

func CreateMigrate(filename string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "internal/migrations", "-seq", filename)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("💥 Create fail <%s>", err)
	}
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
	fmt.Println("✨ Success migrate up")
	if err != nil {
		fmt.Printf("⚡️ migrations up fail <%s>", err)
	}
}

func MigrateSteps(step int) {
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
	err = m.Steps(step)
	fmt.Println("✨ Success migrate up")
	if err != nil {
		fmt.Printf("⚡️ migrations up fail <%s>", err)
	}
}
