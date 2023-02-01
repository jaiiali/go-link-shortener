package link

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jaiiali/go-link-shortener/internal/core/ports"
)

type Handler struct {
	uc       ports.LinkUseCase
	validate *validator.Validate
}

func NewHandler(uc ports.LinkUseCase, router fiber.Router) *Handler {
	handler := &Handler{
		uc:       uc,
		validate: validator.New(),
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
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err := h.validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.uc.Create(req.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var resp = &linkResp{}
	resp.bind(result)

	return c.JSON(resp)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	result, err := h.uc.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var resp = linkListResp{}
	resp.bind(result)

	return c.JSON(resp)
}

func (h *Handler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.uc.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	var resp = &linkResp{}
	resp.bind(result)

	return c.JSON(resp)
}
