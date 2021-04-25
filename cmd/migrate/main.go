package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var userStatement = `create table if not exists "users" (
	"id" integer not null primary key autoincrement,
	"email" varchar not null,
	"password" varchar not null,
	"created_at" datetime null,
	"active" boolean not null check (active in (0, 1)) default 0
);`

var notesStatement = `create table if not exists "notes" (
	"id" integer not null primary key autoincrement,
	"user_id" integer not null ,
	"date" datetime not null default "",
	"updated_at" datetime null,
	"content" text not null default "[]",
	foreign key(user_id) references users(id)
);`

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load env file.")
	}

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(userStatement)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(notesStatement)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrated successfully.")
}

func openDB() (*sql.DB, error) {
	loc := fmt.Sprintf("./database/%s", os.Getenv("DB_NAME"))

	db, err := sql.Open("sqlite3", loc)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
