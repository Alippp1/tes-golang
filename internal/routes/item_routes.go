package routes

import (
	"github.com/Alippp1/tes-golang/internal/database"
	"github.com/Alippp1/tes-golang/internal/handler"
	"github.com/Alippp1/tes-golang/internal/middleware"
	"github.com/Alippp1/tes-golang/internal/repository"
	"github.com/Alippp1/tes-golang/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterItemRoutes(api fiber.Router) {
	itemRepo := repository.NewItemRepository(database.DB)
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	item := api.Group("/item", middleware.AuthRequired())
	item.Post("/", itemHandler.Create)
	item.Get("/", itemHandler.FindAll)
	item.Get("/:id", itemHandler.FindByID)
	item.Put("/:id", itemHandler.Update)
	item.Delete("/:id", middleware.RoleRequired("admin"), itemHandler.Delete)
}
