package handler

import "github.com/gofiber/fiber/v2"

type AnimeHandler struct {
	*Handler
}

func NewAnimeHandler(handler *Handler) *AnimeHandler {
	return &AnimeHandler{
		Handler: handler,
	}
}

func (h *AnimeHandler) Create(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *AnimeHandler) List(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *AnimeHandler) Get(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *AnimeHandler) Update(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *AnimeHandler) Delete(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}
