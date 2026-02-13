package main

import (
	"log"
	"os"

	"github.com/SKjustSK/url-shortner-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/shorten", routes.ShotenURL)
}

func main() {
	// 1. Load .env file (Local dev only)
	// On Render, this does nothing because the file won't exist, which is fine.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := fiber.New()
	app.Use(logger.New())

	// 2. CORS: Strict checking for your Frontend
	// If FRONTEND_DOMAIN is empty, this might block requests.
	frontend := os.Getenv("FRONTEND_DOMAIN")
	if frontend == "" {
		log.Println("Warning: FRONTEND_DOMAIN is not set")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: frontend,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	// 3. Port Handling (Crucial for Render)
	// Render injects "PORT". If not found, fall back to "APP_PORT" or "3000"
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "3000" // Default fallback
	}

	// fiber.Listen expects ":3000", Render gives "10000"
	// We ensure the colon is present.
	log.Fatal(app.Listen(":" + port))
}
