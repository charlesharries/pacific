package mock

import (
	"time"

	"github.com/charlesharries/pacific/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Email:   "alice@example.com",
	Created: time.Now(),
	Active:  true,
}

type UserModel struct{}

// Insert mocks a database insert.
func (m *UserModel) Insert(email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// Authenticate mocks authenticating a user.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "alice@example.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get mocks retrieving a user from the database.
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
