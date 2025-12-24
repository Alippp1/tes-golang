package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Alippp1/tes-golang/internal/database"
	"github.com/Alippp1/tes-golang/internal/handler"
	"github.com/Alippp1/tes-golang/internal/repository"
	"github.com/Alippp1/tes-golang/internal/service"
)

func RegisterAuthRoutes(api fiber.Router) {
	userRepo := repository.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}
