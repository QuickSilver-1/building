package presentation

import (
	_ "building/docs"
	"building/internal/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Структура BuildingServ определяет сервис для работы с зданиями
type BuildingServ struct{
    db *sql.DB
}

// Конструктор для создания нового сервиса BuildingServ
func NewBuildingService(ip, port, user, pass, nameDB string) (*BuildingServ, error) {
    sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", ip, port, user, pass, nameDB)
    conn, err := sql.Open("postgres", sqlInfo)

    if err != nil {
        return nil, fmt.Errorf("failed connection: %v", err)
    }

    return &BuildingServ{
        db: conn,
    }, nil
}

// Метод для создания нового здания
func (b *BuildingServ) Create(building domain.Building) error {
    _, err := b.db.Exec(` INSERT INTO building ("name", "city", "year", "floor") VALUES ($1, $2, $3, $4) `, building.Name, building.City, building.Year, building.Floor)    
    LoggerService.Debug("building has been created")

    if err != nil {
        LoggerService.Warn(fmt.Sprintf("Building created error: %v", err))
        return err
    }

    return nil
}

// Метод для получения зданий с возможностью фильтрации
func (b *BuildingServ) Get(building domain.Building) (*[]domain.Building, error) {
    var rows *sql.Rows
    var err error

    if building.Id != 0 {
        rows, err = b.db.Query(` SELECT * FROM building WHERE "id" = $1 `, building.Id)

        if err != nil {
            LoggerService.Warn(fmt.Sprintf("Getting building has been failed: %v", err))
            return nil, err
        }

        LoggerService.Debug("Getting building has been success")
    } else {
        filter := ""

        if building.Year != "" {
            filter += fmt.Sprintf(`AND "year" = %s`, building.Year)
        }

        if building.Floor != "" {
            filter += fmt.Sprintf(`AND "floor" = %s`, building.Floor)
        }

        building.Name = "%" + building.Name + "%"
        building.City = "%" + building.City + "%"

        rows, err = b.db.Query(` SELECT * FROM building WHERE "name" ILIKE $1 AND "city" ILIKE $2 ` + filter, building.Name, building.City)

        if err != nil {
            LoggerService.Warn(fmt.Sprintf("Getting building has been failed: %v", err))
            return nil, err
        }

        LoggerService.Debug("Getting building has been success")
    }

    var answer []domain.Building
    var id uint64
    var name, city, year, floor string
    for rows.Next() {
        err = rows.Scan(&id, &name, &city, &year, &floor)

        if err != nil {
            return nil, err
        }

        build := domain.Building{
            Id: id,
            Name: name,
            City: city,
            Year: year,
            Floor: floor,
        }

        answer = append(answer, build)
    }

    return &answer, nil
}