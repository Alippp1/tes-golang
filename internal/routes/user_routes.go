package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Alippp1/tes-golang/internal/database"
	"github.com/Alippp1/tes-golang/internal/handler"
	"github.com/Alippp1/tes-golang/internal/middleware"
	"github.com/Alippp1/tes-golang/internal/repository"
	"github.com/Alippp1/tes-golang/internal/service"
)

func RegisterUserRoutes(api fiber.Router) {
	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	users := api.Group("/users", middleware.AuthRequired())
	users.Get("/", userHandler.FindAll)
	users.Get("/:id", userHandler.FindByID)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", middleware.RoleRequired("admin"), userHandler.Delete)
}