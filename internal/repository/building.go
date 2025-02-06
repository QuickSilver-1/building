package repository

import "building/internal/domain"

// Интерфейс BuildingRepo определяет методы для работы с объектами Building
type BuildingRepo interface {
    // Метод для получения списка объектов Building
    Get(domain.Building) (*[]domain.Building, error)
    
    // Метод для создания нового объекта Building
    Create(domain.Building) error
}
