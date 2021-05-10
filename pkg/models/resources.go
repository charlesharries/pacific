package models

import (
	"time"
)

type UserResource struct {
	ID      uint      `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created_at"`
	Active  bool      `json:"active"`
}

type NoteResource struct {
	ID        uint         `json:"id"`
	User      UserResource `json:"user"`
	Date      time.Time    `json:"date"`
	UpdatedAt time.Time    `json:"updated_at"`
	Content   string       `json:"content"`
}

type TodoResource struct {
	ID        uint         `json:"id"`
	User      UserResource `json:"user"`
	Date      time.Time    `json:"date"`
	Completed bool         `json:"completed"`
	Content   string       `json:"content"`
}
