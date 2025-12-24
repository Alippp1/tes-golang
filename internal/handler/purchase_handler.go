package handler

import (
	"github.com/Alippp1/tes-golang/internal/dto"
	"github.com/Alippp1/tes-golang/internal/service"
	"github.com/gofiber/fiber/v2"
)

type PurchasingHandler struct {
	service service.PurchasingService
}

func NewPurchasingHandler(service service.PurchasingService) *PurchasingHandler {
	return &PurchasingHandler{service}
}

func (h *PurchasingHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.CreatePurchasingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	data, err := h.service.Create(userID, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "purchasing created successfully",
		"data":    data,
	})
}

func (h *PurchasingHandler) FindAll(c *fiber.Ctx) error {
	data, err := h.service.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "purchasing list retrieved successfully",
		"data":    data,
	})
}

func (h *PurchasingHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid purchasing id",
		})
	}

	data, err := h.service.FindByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "purchasing retrieved successfully",
		"data":    data,
	})
}

