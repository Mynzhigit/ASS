package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/flowerstore/internal/entity"
	"github.com/yourusername/flowerstore/internal/repository"
)

func TestGetAdminByID(t *testing.T) {
	db, err := database.ConnectDB()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewAdminRepository(db)

	t.Run("Existing admin", func(t *testing.T) {
		admin, err := repo.GetAdminByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, admin)
		assert.Equal(t, 1, admin.ID)
		assert.Equal(t, "testadmin", admin.Username)
		assert.Equal(t, "admin@example.com", admin.Email)
	})

	t.Run("Non-existing admin", func(t *testing.T) {
		admin, err := repo.GetAdminByID(2)
		assert.Error(t, err)
		assert.Nil(t, admin)
	})
}
