package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/flowerstore/internal/entity"
	"github.com/yourusername/flowerstore/internal/repository"
)

func TestGetUserByID(t *testing.T) {
	db, err := database.ConnectDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	t.Run("Existing user", func(t *testing.T) {
		user, err := repo.GetUserByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "testuser", user.Username)
		assert.Equal(t, "test@example.com", user.Email)
	})

	t.Run("Non-existing user", func(t *testing.T) {
		user, err := repo.GetUserByID(2)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
