package presentation

import (
	"building/internal/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Структура Handlers определяет хендлеры для обработки HTTP-запросов
type Handlers struct{}

// Конструктор для создания нового Handlers
func NewHandlers() *Handlers {
    return &Handlers{}
}

// @Summary Create a new building
// @Description Create a new building with the provided details
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body domain.Building true "Building"
// @Success 200 {string} string "The building has been successfully created"
// @Failure 400 {string} string "Invalid input"
// @Router /building [post]
func (*Handlers) CreateBuilding(ctx *gin.Context) { // Cоздания building
    var build domain.Building
    json.NewDecoder(ctx.Request.Body).Decode(&build)

    if _, err := strconv.Atoi(build.Year); err != nil && build.Year != "" {
        ctx.JSON(http.StatusBadRequest, "Invalid year")
        return
    }

    if _, err := strconv.Atoi(build.Floor); err != nil && build.Floor != "" {
        ctx.JSON(http.StatusBadRequest, "Invalid floor")
        return
    }

    err := BuildingService.Create(build)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

    ctx.JSON(http.StatusOK, "The building has been successfully created")
}

// @Summary Get buildings
// @Description Get buildings by filters
// @Tags buildings
// @Accept json
// @Produce json
// @Param id query string false "Building ID"
// @Param name query string false "Building name"
// @Param city query string false "Building city"
// @Param year query string false "Year of construction"
// @Param floor query string false "Number of floors"
// @Success 200 {array} domain.Building
// @Failure 400 {string} string "Invalid input"
// @Router /building [get]
func (*Handlers) GetBuilding(ctx *gin.Context) { // Получения building
    idStr := ctx.Request.URL.Query().Get("id")
    name := ctx.Request.URL.Query().Get("name")
    city := ctx.Request.URL.Query().Get("city")
    year := ctx.Request.URL.Query().Get("year")
    floor := ctx.Request.URL.Query().Get("floor")

    if idStr == "" {
        idStr = "0"
    }
    
    id, err := strconv.ParseUint(idStr, 10, 64)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

    if _, err := strconv.Atoi(year); err != nil && year != "" {
        ctx.JSON(http.StatusBadRequest, "Invalid year")
        return
    }

    if _, err := strconv.Atoi(floor); err != nil && floor != "" {
        ctx.JSON(http.StatusBadRequest, "Invalid floor")
        return
    }

    building := domain.Building{
        Id: id,
        Name: name,
        City: city,
        Year: year,
        Floor: floor,
    }

    buildings, err := BuildingService.Get(building)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }

    if len(*buildings) == 0 {
        ctx.JSON(http.StatusBadRequest, "There are no buildings with such filters")
        return
    }

    ctx.JSON(http.StatusOK, buildings)
}

func (*Handlers) GetSwagger(ctx *gin.Context) {
        jsonFile, err := os.Open("../../../docs/swagger.json")

        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open JSON file"})
            return
        }

        defer jsonFile.Close()
        byteValue, err := ioutil.ReadAll(jsonFile)
        
		if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
            return
        }

        var swagger interface{} 
        json.Unmarshal(byteValue, &swagger)

        ctx.JSON(http.StatusOK, swagger)
    }