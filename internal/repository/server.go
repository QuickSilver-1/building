package repository

// Интерфейс ServerRepo определяет методы для управления сервером
type ServerRepo interface {
	Start() error
}