package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/flowerstore/internal/delivery"
	"github.com/yourusername/flowerstore/internal/repository"
	"github.com/yourusername/flowerstore/internal/usecase"
)

func TestIntegration(t *testing.T) {
	// Подготовка объектов, необходимых для работы теста
	db, err := database.ConnectDB()
	assert.NoError(t, err)
	defer db.Close()

	// Инициализация репозиториев и бизнес-логики
	flowerRepo := repository.NewFlowerRepository(db)
	catalogUsecase := usecase.NewCatalogUsecase(flowerRepo)

	// Создание HTTP обработчиков
	handler := delivery.NewHandler(catalogUsecase)

	// Создание фейкового HTTP сервера
	server := httptest.NewServer(http.HandlerFunc(handler.CatalogHandler))
	defer server.Close()

	// Создание клиента для отправки запросов к серверу
	client := &http.Client{}

	// Создание запроса к API
	req, err := http.NewRequest("GET", server.URL+"/catalog", nil)
	assert.NoError(t, err)

	// Отправка запроса к API
	res, err := client.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Проверка содержимого ответа (может быть дополнена)
	// Например, можно проверить, что тело ответа содержит ожидаемый JSON
}
