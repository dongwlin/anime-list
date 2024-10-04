package handler

import (
	"context"
	"time"

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

type createSeasonRequest struct {
	Name        string    `json:"name"`
	Value       int64     `json:"value"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
}

func (sh *SeasonHandler) Create(c *fiber.Ctx) error {
	animeId, err := c.ParamsInt("animeId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &createSeasonRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	params := operator.CreateSeasonParams{
		AnimeID:     animeId,
		Name:        req.Name,
		Value:       req.Value,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	season, err := sh.operator.Create(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", season, nil))
}

type listSeasonRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (sh *SeasonHandler) List(c *fiber.Ctx) error {
	animeId, err := c.ParamsInt("animeId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &listSeasonRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	params := operator.ListSeasonParams{
		AnimeID: animeId,
		Offset:  (req.Page - 1) * req.Size,
		Limit:   req.Size,
	}
	ctx := context.Background()
	total, seasons, err := sh.operator.List(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	totalPages := total / req.Size
	if total%req.Size != 0 {
		totalPages += 1
	}
	pagnation := Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		PerPage:     req.Size,
		CurrentPage: req.Page,
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", seasons, pagnation))
}

type searchSeasonRequest struct {
	Query string `query:"query"`
	Page  int    `query:"page"`
	Size  int    `query:"size"`
}

func (sh *SeasonHandler) Search(c *fiber.Ctx) error {
	animeId, err := c.ParamsInt("animeId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &searchSeasonRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	params := operator.SearchSeasonParams{
		Query:   req.Query,
		AnimeID: animeId,
		Offset:  (req.Page - 1) * req.Size,
		Limit:   req.Size,
	}
	ctx := context.Background()
	total, seasons, err := sh.operator.Search(ctx, params)
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
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", seasons, pagination))
}

func (sh *SeasonHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	ctx := context.Background()
	season, err := sh.operator.Get(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", season, nil))
}

type updateSeasonRequest struct {
	Name        string    `json:"name"`
	Value       int64     `json:"value"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
}

func (sh *SeasonHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &updateSeasonRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	params := operator.UpdateSeasonParams{
		ID:          id,
		Name:        req.Name,
		Value:       req.Value,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	season, err := sh.operator.Update(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", season, nil))
}

func (sh *SeasonHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	ctx := context.Background()
	err = sh.operator.Delete(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", nil, nil))
}
