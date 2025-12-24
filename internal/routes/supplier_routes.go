package routes

import (
	"github.com/Alippp1/tes-golang/internal/database"
	"github.com/Alippp1/tes-golang/internal/handler"
	"github.com/Alippp1/tes-golang/internal/middleware"
	"github.com/Alippp1/tes-golang/internal/repository"
	"github.com/Alippp1/tes-golang/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterSupplierRoutes(api fiber.Router) {
	supplierRepo := repository.NewSupplierRepository(database.DB)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	supplier := api.Group("/supplier", middleware.AuthRequired())
	supplier.Post("/", supplierHandler.Create)
	supplier.Get("/", supplierHandler.FindAll)
	supplier.Get("/:id", supplierHandler.FindByID)
	supplier.Put("/:id", supplierHandler.Update)
	supplier.Delete("/:id", middleware.RoleRequired("admin"), supplierHandler.Delete)

}
