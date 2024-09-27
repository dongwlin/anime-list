package handler

import "github.com/gofiber/fiber/v2"

type PingHandler struct {
	*Handler
}

func NewPingHandler(handler *Handler) *PingHandler {
	return &PingHandler{
		Handler: handler,
	}
}

func (h *PingHandler) Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("pong!!!!!")
}
