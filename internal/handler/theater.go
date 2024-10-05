package handler

import (
	"context"
	"time"

	"github.com/dongwlin/anime-list/internal/operator"
	"github.com/gofiber/fiber/v2"
)

type TheaterHandler struct {
	operator operator.TheaterOperator
}

func NewTheaterHandler(operator operator.TheaterOperator) *TheaterHandler {
	return &TheaterHandler{
		operator: operator,
	}
}

type createTheaterRequest struct {
	Name        string    `json:"name"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
}

func (th *TheaterHandler) Create(c *fiber.Ctx) error {
	animeId, err := c.ParamsInt("animeId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &createTheaterRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	params := operator.CreateTheaterParams{
		AnimeID:     animeId,
		Name:        req.Name,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	theater, err := th.operator.Create(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", theater, nil))
}

type listTheaterRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (th *TheaterHandler) List(c *fiber.Ctx) error {
	animeId, err := c.ParamsInt("animeId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &listTheaterRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	params := operator.ListTheaterParams{
		AnimeID: animeId,
		Offset:  (req.Page - 1) * req.Size,
		Limit:   req.Size,
	}
	ctx := context.Background()
	total, theaters, err := th.operator.List(ctx, params)
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
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", theaters, pagination))
}

type searchTheaterRequest struct {
	Query string `query:"query"`
	Page  int    `query:"page"`
	Size  int    `query:"size"`
}

func (th *TheaterHandler) Search(c *fiber.Ctx) error {
	animeId, err := c.ParamsInt("animeId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &searchTheaterRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	params := operator.SearchTheaterParams{
		AnimeID: animeId,
		Query:   req.Query,
		Offset:  (req.Page - 1) * req.Size,
		Limit:   req.Size,
	}
	ctx := context.Background()
	total, theaters, err := th.operator.Search(ctx, params)
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
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", theaters, pagination))
}

func (th *TheaterHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params", err.Error())
	}
	ctx := context.Background()
	theater, err := th.operator.Get(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", theater, nil))
}

type updateTheaterRequest struct {
	Name        string    `json:"name"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
}

func (th *TheaterHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &updateTheaterRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	params := operator.UpdateTheaterParams{
		ID:          id,
		Name:        req.Name,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	theater, err := th.operator.Update(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", theater, nil))
}

func (th *TheaterHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	ctx := context.Background()
	err = th.operator.Delete(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", nil, nil))
}
