package health

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func NewHandler(router fiber.Router) *Handler {
	handler := &Handler{}

	router.Get("/health", handler.Health)

	return handler
}

func (h *Handler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
