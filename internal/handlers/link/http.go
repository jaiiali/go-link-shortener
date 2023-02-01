package link

import (
	"errors"

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

	router.Post("/links", handler.Create)
	router.Get("/links", handler.FindAll)
	router.Get("/links/:id", handler.FindByID)

	return handler
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req = &linkReq{}

	c.Request().Header.Add("Content-Type", "application/json")
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if req.URL == "" {
		return errors.New("url is required")
	}

	result, err := h.uc.Create(req.URL)
	if err != nil {
		return err
	}

	var resp = &linkResp{}
	resp.bind(result)

	return c.JSON(resp)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	result, err := h.uc.FindAll()
	if err != nil {
		return err
	}

	var resp = linkListResp{}
	resp.bind(result)

	return c.JSON(resp)
}

func (h *Handler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.uc.FindByID(id)
	if err != nil {
		return err
	}

	var resp = &linkResp{}
	resp.bind(result)

	return c.JSON(resp)
}
