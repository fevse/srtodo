package main

import (
	"fmt"
	"log"

	"github.com/fevse/srtodo/internal/conf"
	"github.com/fevse/srtodo/internal/server"
	"github.com/fevse/srtodo/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	dbConf, err := conf.DbConf()
	if err != nil {
		log.Fatalf("DB configuration error: %v", err)
	}

	serverConf, err := conf.ServerConf()
	if err != nil {
		log.Fatalf("Server configuration error: %v", err)
	}

	err = storage.NewStorage(dbConf)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer storage.Storage.Close()

	fmt.Println("Storage: done")

	err = storage.Migrate(".")

	if err != nil {
		log.Fatalf("DB migration error: %v", err)
	}

	fmt.Println("Migrate: done")

	app := fiber.New()

	fmt.Println("Server: done")

	app.Use(logger.New())

	server.Routes(app)

	log.Fatal(app.Listen(serverConf))
}
