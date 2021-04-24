package models

import (
	"errors"
	"time"
)

// ErrNoRecord displays if a user requests a non-existent resource.
var ErrNoRecord = errors.New("models: no matching record found")

// ErrInvalidCredentials displays if a user tries to log in with an
// incorrect email address or password.
var ErrInvalidCredentials = errors.New("models: invalid credentials")

// ErrDuplicateEmail displays when a user tries to sign up with an
// existing email address.
var ErrDuplicateEmail = errors.New("models: duplicate email")

type User struct {
	ID       int
	Email    string
	Password []byte
	Created  time.Time
	Active   bool
}
