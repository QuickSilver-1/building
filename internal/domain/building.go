package domain

// Структура Building описывает здание с различными атрибутами
type Building struct {
    Id    uint64 `json:"id"`    // Уникальный идентификатор здания
    Name  string `json:"name"`  // Название здания
    City  string `json:"city"`  // Город, где находится здание
    Year  string `json:"year"`  // Год постройки
    Floor string `json:"floor"` // Количество этажей
}

// Конструктор для создания новой структуры Building
func NewBuilding(id uint64, name, city, year, floor string) *Building {
    return &Building{
        Id:    id,
        Name:  name,
        City:  city,
        Year:  year,
        Floor: floor,
    }
}
