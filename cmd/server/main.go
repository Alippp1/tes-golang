package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/Alippp1/tes-golang/internal/database"
	"github.com/Alippp1/tes-golang/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}

	app := fiber.New()

	// ✅ CORS WAJIB DI SINI
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	database.ConnectDatabase()
	database.AutoMigrate()

	routes.RegisterRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("🚀 Server running at http://localhost:" + port)
	log.Fatal(app.Listen(":" + port))
}
