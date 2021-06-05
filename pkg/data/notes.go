package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/charlesharries/pacific/pkg/validator"
)

type NoteModel struct {
	DB *sql.DB
}

type Note struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - directive hides field from JSON
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
	UserID    int64     `json:"user_id"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Version   int32     `json:"-"`
}

// Get fetches a specific movie record from the database.
func (m NoteModel) Get(date time.Time, userID int64) (*Note, error) {
	if userID < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, content
		FROM notes
		WHERE user_id = ? AND date = ?`

	note := Note{
		Date:   date,
		UserID: userID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, userID, date).Scan(
		&note.ID,
		&note.CreatedAt,
		&note.Content,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &note, nil
}

// Update changes a specific movie record in the database.
func (m NoteModel) Upsert(note *Note) error {
	query := `
		INSERT INTO notes(user_id, date, content, created_at, updated_at)
		VALUES (?, ?, ?, datetime('now'), datetime('now'))
		ON CONFLICT(date, user_id) DO UPDATE SET content = ?, updated_at = datetime('now')
		RETURNING id, version`

	args := []interface{}{
		note.UserID,
		note.Date,
		note.Content,
		note.Content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&note.ID,
		&note.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	return nil
}

// ValidateNote runs all of our validation checks on the given Movie.
func ValidateNote(v *validator.Validator, note *Note) {
	// TODO
}
