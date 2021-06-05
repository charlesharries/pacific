package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrEditConflict       = errors.New("edit conflict")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
)

// create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and a PermissionModel, as our build progresses.
type Models struct {
	Users UserModel
	Notes NoteModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialised MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Users: UserModel{DB: db},
		Notes: NoteModel{DB: db},
	}
}
