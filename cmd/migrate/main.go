package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://./migrations",
		"sqlite3://database/new.sqlite",
	)

	m.Log = Logger{}

	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("pass an argument: 'up' or 'down'")
	}

	dir := os.Args[1]

	switch dir {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		log.Fatal("you gotta pass 'up' or 'down'")
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("migrated.")
}
