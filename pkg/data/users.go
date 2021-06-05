package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/charlesharries/pacific/pkg/validator"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - directive hides field from JSON
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Active    bool      `json:"active"`
	Version   int32     `json:"-"`
}

// Insert adds a new user
func (m *UserModel) Insert(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (email, password, active, version, created_at)
		VALUES(?, ?, FALSE, 1, datetime('now'))`

	_, err = m.DB.Exec(stmt, email, string(hashedPassword))
	var sqlErr sqlite3.Error
	if err != nil {
		if errors.As(err, &sqlErr) && sqlErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return ErrDuplicateEmail
		}

		return err
	}

	return nil
}

// Get fetches a specific movie record from the database.
func (m UserModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, email, password, active, version
		FROM users
		WHERE id = ?`

	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Email,
		&user.Password,
		&user.Active,
		&user.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

// Update changes a specific movie record in the database.
func (m UserModel) Update(user *User) error {
	query := `
		UPDATE users
		SET email = ?, password = ? version = version + 1
		WHERE id = ? AND version = ?
		RETURNING version`

	args := []interface{}{
		user.Email,
		user.Password,
		user.ID,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

// Delete removes a specific movie record from the database.
func (m UserModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM users
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// Authenticate verifies that a user exists with the provided email
// and password. Returns the relevant user ID if found.
func (m *UserModel) Authenticate(email, password string) (int64, error) {
	var id int64
	var hashedPassword []byte

	// Get the user from the db
	stmt := `SELECT id, password FROM users WHERE email = ? AND active = TRUE`
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		}

		return 0, err
	}

	// Check the password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		}

		return 0, err
	}

	return id, nil
}

// ValidateUser runs all of our validation checks on the given Movie.
func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) <= 500, "email", "must not be more than 500 bytes long")
	v.Check(validator.Matches(user.Email, validator.EmailRX), "email", "must be a valid")

	v.Check(user.Password != "", "password", "must be provided")
	v.Check(len(user.Password) <= 500, "password", "must not be more than 500 bytes long")
}
