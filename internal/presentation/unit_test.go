package presentation

import (
	"building/internal/domain"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Функция для создания тестового контекста
func createTestContext(method, path string, body interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	LoggerService = NewLogger()

	err := godotenv.Load("../../.env")

	if err != nil {
		LoggerService.Fatal(err.Error())
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	BuildingService, err = NewBuildingService(host, port, user, password, name)

	if err != nil {
		LoggerService.Fatal(err.Error())
	}

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    jsonBody, _ := json.Marshal(body)
    req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    c.Request = req

    return c, w
}

// Тест для эндпоинта CreateBuilding
func TestCreateBuilding(t *testing.T) {

    // Инициализация обработчиков и тестового контекста
    handlers := NewHandlers()
    body := domain.Building{
        Name:  "Test Building",
        City:  "Test City",
        Year:  "2022",
        Floor: "10",
    }
    c, w := createTestContext("POST", "/building", body)

    // Запуск обработчика
    handlers.CreateBuilding(c)

    // Проверка ответа
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `"The building has been successfully created"`, w.Body.String())
}

// Тест для эндпоинта GetBuilding
func TestGetBuilding(t *testing.T) {
    // Инициализация сервиса и логгера
    buildingService := &BuildingServ{}
    loggerService := &Logger{}
    BuildingService = buildingService
    LoggerService = loggerService

    // Инициализация обработчиков и тестового контекста
    handlers := NewHandlers()
    query := "id=1&name=Test Building&city=Test City"
    c, w := createTestContext("GET", "/building?"+query, nil)

    // Запуск обработчика
    handlers.GetBuilding(c)

    // Проверка ответа
    assert.Equal(t, http.StatusOK, w.Code)
    // Здесь вы можете добавить более точные проверки в зависимости от ожидаемого результата
}