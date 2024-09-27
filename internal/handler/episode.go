package handler

import "github.com/gofiber/fiber/v2"

type EpisodeHandler struct {
	*Handler
}

func NewEpisodeHandler(handler *Handler) *EpisodeHandler {
	return &EpisodeHandler{
		Handler: handler,
	}
}

func (h *EpisodeHandler) Create(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *EpisodeHandler) List(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *EpisodeHandler) Get(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *EpisodeHandler) Update(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *EpisodeHandler) Delete(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}
