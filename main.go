package main

import (
	"log"

	"github.com/fierzahaikkal/neocourse-be/db"
	"github.com/fierzahaikkal/neocourse-be/internal/configs"
	"github.com/fierzahaikkal/neocourse-be/pkg/utils"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main(){
	config := configs.LoadConfig()
	logger := utils.NewLogger()

	// Database connection
	dbConn := configs.InitDB(config)

	// Run migrations
	if err := db.Migrate(dbConn); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// apply fiber
	app := fiber.New()

	//apply cors
	app.Use(cors.New())

	// will be available at /api/v1/docs
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./api/v1/api-spec.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))

	//apply recover middleware
	app.Use(recover.New())

	//repositories

	//use case

	//route

	// -- auth

	// -- book

	// Listen Server
	logger.Info("Server started on http://localhost:8081")
	if err := app.Listen("localhost:8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}