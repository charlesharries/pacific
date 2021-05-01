package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/charlesharries/pacific/pkg/models"
	mysqlite "github.com/charlesharries/pacific/pkg/models/sqlite"
	"github.com/golangcollege/sessions"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	templateCache map[string]*template.Template
	gorm          *gorm.DB
	users         interface {
		Insert(string, string) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load env file.")
	}

	// If we're in production, we want to save the time that the
	// server starts, to version our assets.
	if os.Getenv("ENV") == "production" {
		os.Setenv("CACHE_KEY", fmt.Sprint(time.Now().Unix()))
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	gorm, err := openGorm()
	if err != nil {
		errorLog.Fatal(err)
	}

	gorm.AutoMigrate(&models.Note{})

	templateCache, err := newTemplateCache("./resources/views")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(os.Getenv("APP_SECRET")))
	session.Lifetime = 24 * time.Hour
	gob.Register(TemplateUser{})

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		templateCache: templateCache,
		users:         &mysqlite.UserModel{DB: db},
		gorm:          gorm,
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server at http://%s\n", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openGorm() (*gorm.DB, error) {
	loc := fmt.Sprintf("./database/%s", os.Getenv("DB_NAME"))

	return gorm.Open(sqlite.Open(loc), &gorm.Config{})
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
