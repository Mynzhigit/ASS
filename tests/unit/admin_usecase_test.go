package unit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/flowerstore/internal/entity"
	"github.com/yourusername/flowerstore/internal/usecase"
)

type mockAdminRepository struct{}

func (m *mockAdminRepository) GetAdminByID(adminID int) (*entity.Admin, error) {
	if adminID == 1 {
		return &entity.Admin{ID: 1, Username: "testadmin", Email: "admin@example.com"}, nil
	}
	return nil, errors.New("admin not found")
}

func TestGetAdminByID(t *testing.T) {
	repo := &mockAdminRepository{}
	usecase := usecase.NewAdminUsecase(repo)

	t.Run("Existing admin", func(t *testing.T) {
		admin, err := usecase.GetAdminByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, admin)
		assert.Equal(t, 1, admin.ID)
		assert.Equal(t, "testadmin", admin.Username)
		assert.Equal(t, "admin@example.com", admin.Email)
	})

	t.Run("Non-existing admin", func(t *testing.T) {
		admin, err := usecase.GetAdminByID(2)
		assert.Error(t, err)
		assert.Nil(t, admin)
		assert.EqualError(t, err, "admin not found")
	})
}
