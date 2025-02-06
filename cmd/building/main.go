// @title Building API
// @version 1.0.0
// @description Тестовое задание для компании leadgen
// @host 89.46.131.181:8081
package main

import (
	"building/internal/presentation"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := presentation.NewLogger()

	err := godotenv.Load("../../../.env")

	if err != nil {
		logger.Fatal(err.Error())
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	building, err := presentation.NewBuildingService(host, port, user, password, name)

	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	server := presentation.NewServer(building, logger)
	err = server.Start()

	if err != nil {
		logger.Fatal(err.Error())
	}
}