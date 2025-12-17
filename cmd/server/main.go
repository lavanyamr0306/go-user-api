package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/lavanyamr0306/go-user-api/internal/db/sqlc" // SQLC package
	"github.com/lavanyamr0306/go-user-api/internal/handler"
	"github.com/lavanyamr0306/go-user-api/internal/routes"
	"github.com/lavanyamr0306/go-user-api/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Database connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		sugar.Fatalw("cannot connect to db", "error", err)
	}
	defer conn.Close()

	// SQLC queries
	queries := sqlc.New(conn)

	// Services & Handlers
	userService := service.NewUserService(queries)
	userHandler := handler.NewUserHandler(userService)

	// Fiber app
	app := fiber.New()

	// Routes
	routes.RegisterRoutes(app, userHandler)

	// Health endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	sugar.Infow("server started", "port", port)
	log.Fatal(app.Listen(":" + port))
}
