package mock

import (
	"time"

	"github.com/mgeale/homeserver/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Alice",
	Email:   "alice@example.com",
	Created: time.Now(),
	Active:  true,
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) error {
	switch email {
	case "alice@example.com":
		return nil
	default:
		return models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}