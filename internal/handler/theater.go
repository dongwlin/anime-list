package handler

import "github.com/gofiber/fiber/v2"

type TheaterHandler struct {
	*Handler
}

func NewTheaterHandler(handler *Handler) *TheaterHandler {
	return &TheaterHandler{
		Handler: handler,
	}
}

func (h *TheaterHandler) Create(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *TheaterHandler) List(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *TheaterHandler) Get(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *TheaterHandler) Update(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func (h *TheaterHandler) Delete(c *fiber.Ctx) error {
	// todo
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}
