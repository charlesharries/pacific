package sqlite

import (
	"database/sql"
	"errors"
	"time"

	"github.com/charlesharries/pacific/pkg/models"
)

type NoteModel struct {
	DB *sql.DB
}

// Insert adds a new user
func (m *NoteModel) Insert(userID int, date time.Time, content string) error {
	stmt := `INSERT INTO notes (user_id, date, updated_at, content) VALUES(?, ?, datetime(now), ?)`

	_, err := m.DB.Exec(stmt, userID, date.String(), content)
	if err != nil {
		return err
	}

	return nil
}

// Get fetches a user from the database.
func (m *NoteModel) Get(id int) (*models.Note, error) {
	n := &models.Note{}

	stmt := `SELECT id, user_id, date, updated_at, content FROM notes WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&n.ID, &n.UserID, &n.Date, &n.UpdatedAt, &n.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return n, nil
}

// GetAll fetches all notes for the given user.
func (m *NoteModel) GetAll(userID int) ([]*models.Note, error) {
	var notes []*models.Note

	stmt := `select id, user_id, date, updated_at, content FROM notes WHERE user_id = ?`
	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return notes, err
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.Note{}
		err := rows.Scan(&n.ID, &n.UserID, &n.Date, &n.UpdatedAt, &n.Content)
		if err != nil {
			return notes, err
		}
		notes = append(notes, n)
	}

	err = rows.Err()
	if err != nil {
		return notes, err
	}

	return notes, nil
}
