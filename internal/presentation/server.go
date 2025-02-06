package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Переменные для доступа к сервисам
var (
    BuildingService *BuildingServ
    LoggerService   *Logger
)

// Структура Server определяет сервер с сервисами BuildingServ и Logger
type Server struct {
    building    *BuildingServ
    logger      *Logger
    srv         *gin.Engine
}

// Конструктор для создания нового сервера
func NewServer(building *BuildingServ, logger *Logger) *Server {
    gin.SetMode(gin.ReleaseMode)
    srv := gin.New()
    
    h := NewHandlers()
    srv.Use(Middleware())
    
    // Регистрация маршрута Swagger-документации
    srv.GET("/help", h.GetSwagger)

    // Маршруты для создания и получения зданий
    srv.POST("/building", h.CreateBuilding)
    srv.GET("/building", h.GetBuilding)

    // Обработка несуществующих маршрутов
    srv.NoRoute(func(ctx *gin.Context) {
        ctx.JSON(http.StatusBadRequest, "Invalid route")
    })

    logger.Info("Server has been created")
    return &Server{
        srv: srv,
        building: building,
        logger: logger,
    }
}

// Метод для запуска сервера
func (s *Server) Start() error {
    BuildingService = s.building
    LoggerService = s.logger

    LoggerService.Info("Startting server")
    err := s.srv.Run(":8081")

    if err != nil {
        LoggerService.Error("Server error")
        return err
    }

    LoggerService.Info("Stopping server")
    return nil
}