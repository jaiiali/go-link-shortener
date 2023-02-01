package redirect

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaiiali/go-link-shortener/internal/core/ports"
)

type Handler struct {
	uc ports.LinkUseCase
}

func NewHandler(uc ports.LinkUseCase, router fiber.Router) *Handler {
	handler := &Handler{
		uc: uc,
	}

	router.Get("/:id", handler.FindByID)

	return handler
}

func (h *Handler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.uc.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Redirect(result.Original, fiber.StatusTemporaryRedirect)
}
