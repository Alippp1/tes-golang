package routes

import (
	"github.com/Alippp1/tes-golang/internal/database"
	"github.com/Alippp1/tes-golang/internal/handler"
	"github.com/Alippp1/tes-golang/internal/middleware"
	"github.com/Alippp1/tes-golang/internal/repository"
	"github.com/Alippp1/tes-golang/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterPurchaseRoutes(api fiber.Router) {
	purchasingRepo := repository.NewPurchasingRepository(database.DB)
	purchasingDetailRepo := repository.NewPurchasingDetailRepository(database.DB)
	itemRepo := repository.NewItemRepository(database.DB)

	purchasingService := service.NewPurchasingService(
		database.DB,
		purchasingRepo,
		purchasingDetailRepo,
		itemRepo,
	)

	purchasingHandler := handler.NewPurchasingHandler(purchasingService)

	purchase := api.Group("/purchase", middleware.AuthRequired())
	purchase.Post("/", purchasingHandler.Create)
	purchase.Get("/", purchasingHandler.FindAll)          
	purchase.Get("/:id", purchasingHandler.FindByID)      
	// purchase.Put("/:id", purchasingHandler.Update)        
	// purchase.Delete("/:id", purchasingHandler.Delete)     
}