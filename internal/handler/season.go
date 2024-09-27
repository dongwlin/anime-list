package handler

import "github.com/gofiber/fiber/v2"

type SeasonHandler struct {
	*Handler
}

func NewSeasonHandler(handler *Handler) *SeasonHandler {
	return &SeasonHandler{
		Handler: handler,
	}
}

func (h *SeasonHandler) Create(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *SeasonHandler) List(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *SeasonHandler) Get(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *SeasonHandler) Update(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *SeasonHandler) Delete(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}
