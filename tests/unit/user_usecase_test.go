package unit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/flowerstore/internal/entity"
	"github.com/yourusername/flowerstore/internal/usecase"
)

type mockUserRepository struct{}

func (m *mockUserRepository) GetUserByID(userID int) (*entity.User, error) {
	if userID == 1 {
		return &entity.User{ID: 1, Username: "testuser", Email: "test@example.com"}, nil
	}
	return nil, errors.New("user not found")
}

func TestGetUserByID(t *testing.T) {
	repo := &mockUserRepository{}
	usecase := usecase.NewUserUsecase(repo)

	t.Run("Existing user", func(t *testing.T) {
		user, err := usecase.GetUserByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "testuser", user.Username)
		assert.Equal(t, "test@example.com", user.Email)
	})

	t.Run("Non-existing user", func(t *testing.T) {
		user, err := usecase.GetUserByID(2)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})
}
