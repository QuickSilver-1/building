package presentation

import (
	"fmt"

	"go.uber.org/zap"
)

// Структура Logger определяет логгер с использованием zap
type Logger struct {
    logger *zap.Logger
}

// Конструктор для создания нового логгера
func NewLogger() *Logger {
    config := zap.NewDevelopmentConfig()
    config.OutputPaths = []string{"stdout"}

    logger, err := config.Build()

    if err != nil {
        panic(fmt.Sprintf("failed to configure logger: %v", err))
    }

    return &Logger{
        logger: logger,
    }
}

func (l *Logger) Fatal(msg string) {
    l.logger.Fatal(msg)
}

func (l *Logger) Error(msg string) {
    l.logger.Error(msg)
}

func (l *Logger) Warn(msg string) {
    l.logger.Warn(msg)
}

func (l *Logger) Info(msg string) {
    l.logger.Info(msg)
}

func (l *Logger) Debug(msg string) {
    l.logger.Debug(msg)
}
