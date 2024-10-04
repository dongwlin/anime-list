package handler

import (
	"context"

	"github.com/dongwlin/anime-list/internal/operator"
	"github.com/gofiber/fiber/v2"
)

type AnimeHandler struct {
	operator operator.AnimeOperator
}

func NewAnimeHandler(operator operator.AnimeOperator) *AnimeHandler {
	return &AnimeHandler{
		operator: operator,
	}
}

type createAnimeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (ah *AnimeHandler) Create(c *fiber.Ctx) error {
	req := &createAnimeRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}

	params := operator.CreateAnimeParams{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	anime, err := ah.operator.Create(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", anime, nil))
}

type listAnimeRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (ah *AnimeHandler) List(c *fiber.Ctx) error {
	req := &listAnimeRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invaild params", err.Error()))
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	params := operator.ListAnimeParams{
		Offset: (req.Page - 1) * req.Size,
		Limit:  req.Size,
	}

	ctx := context.Background()
	total, animes, err := ah.operator.List(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}

	totalPages := total / req.Size
	if total%req.Size != 0 {
		totalPages += 1
	}

	pagination := Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		PerPage:     req.Size,
		CurrentPage: req.Page,
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", animes, pagination))
}

type searchAnimeRequest struct {
	Query string `query:"query"`
	Page  int    `query:"page"`
	Size  int    `query:"size"`
}

func (ah *AnimeHandler) Search(c *fiber.Ctx) error {
	req := &searchAnimeRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	params := operator.SearchAnimeParams{
		Query:  req.Query,
		Offset: (req.Page - 1) * req.Size,
		Limit:  req.Size,
	}
	ctx := context.Background()
	total, animes, err := ah.operator.Search(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Internal server error", err.Error()))
	}

	totalPages := total / req.Size
	if total%req.Size != 0 {
		totalPages += 1
	}

	pagination := Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		PerPage:     req.Size,
		CurrentPage: req.Page,
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", animes, pagination))
}

func (ah *AnimeHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params", err.Error())
	}
	ctx := context.Background()
	anime, err := ah.operator.Get(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Interanl server error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", anime, nil))
}

type updateAnimeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (ah *AnimeHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params", err.Error())
	}
	req := &updateAnimeRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params", err.Error())
	}

	params := operator.UpdateAnimeParams{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	anime, err := ah.operator.Update(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", anime, nil))
}

func (ah *AnimeHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	ctx := context.Background()
	err = ah.operator.Delete(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", nil, nil))
}
