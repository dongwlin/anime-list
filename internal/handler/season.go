package handler

import (
	"github.com/dongwlin/anime-list/internal/operator"
	"github.com/gofiber/fiber/v2"
)

type SeasonHandler struct {
	operator operator.SeasonOperator
}

func NewSeasonHandler(operator operator.SeasonOperator) *SeasonHandler {
	return &SeasonHandler{
		operator: operator,
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
