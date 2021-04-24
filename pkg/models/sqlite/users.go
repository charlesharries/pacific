package sqlite

import (
	"database/sql"
	"errors"

	"github.com/charlesharries/pacific/pkg/models"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

// Insert adds a new user
func (m *UserModel) Insert(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (email, password, created_at) VALUES(?, ?, datetime('now'))`

	_, err = m.DB.Exec(stmt, email, string(hashedPassword))
	if err != nil {
		if errors.Is(err, sqlite3.ErrConstraintUnique) {
			return models.ErrDuplicateEmail
		}

		return err
	}

	return nil
}

// Authenticate verifies that a user exists with the provided email
// and password. Returns the relevant user ID if found.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	// Get the user from the db
	stmt := `SELECT id, password FROM users WHERE email = ? AND active = TRUE`
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		}

		return 0, err
	}

	// Check the password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}

		return 0, err
	}

	return id, nil
}

// Get fetches a user from the database.
func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}

	stmt := `SELECT id, email, created_at, active FROM users WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&u.ID, &u.Email, &u.Created, &u.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return u, nil
}
