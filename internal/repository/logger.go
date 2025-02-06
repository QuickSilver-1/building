package repository

// Интерфейс LoggerRepo определяет методы для логирования сообщений
type LoggerRepo interface {
	Fatal(string)
	Error(string)
	Warn(string)
	Info(string)
	Debug(string)
}