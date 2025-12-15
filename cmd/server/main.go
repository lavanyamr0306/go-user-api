package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lavanyamr0306/go-user-api/internal/db"
	"github.com/lavanyamr0306/go-user-api/internal/handler"
	"github.com/lavanyamr0306/go-user-api/internal/routes"
	"github.com/lavanyamr0306/go-user-api/internal/service"
)

func main() {
	// Load .env
	godotenv.Load()

	// Connect to DB
	conn := db.Connect()
	defer conn.Close()

	queries := db.New(conn) // SQLC generated constructor

	// Create services and handlers
	userService := service.NewUserService(queries)
	userHandler := handler.NewUserHandler(userService)

	// Fiber app
	app := fiber.New()

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Register user routes
	routes.RegisterRoutes(app, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Fatal(app.Listen(":" + port))
}
