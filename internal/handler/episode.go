package handler

import (
	"context"

	"github.com/dongwlin/anime-list/internal/operator"
	"github.com/gofiber/fiber/v2"
)

type EpisodeHandler struct {
	operator operator.EpisodeOperator
}

func NewEpisodeHandler(operator operator.EpisodeOperator) *EpisodeHandler {
	return &EpisodeHandler{
		operator: operator,
	}
}

type createEpisodeRequest struct {
	Name        string `json:"name"`
	Value       int64  `json:"value"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (eh *EpisodeHandler) Create(c *fiber.Ctx) error {
	seasonId, err := c.ParamsInt("seasonId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &createEpisodeRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	params := operator.CreateEpisodeParams{
		SeasonID:    seasonId,
		Name:        req.Name,
		Value:       req.Value,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	episode, err := eh.operator.Create(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", episode, nil))
}

type listEpisodeRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

func (eh *EpisodeHandler) List(c *fiber.Ctx) error {
	seasonId, err := c.ParamsInt("seasonId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &listEpisodeRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	params := operator.ListEpisodeParams{
		SeasonID: seasonId,
		Offset:   (req.Page - 1) * req.Size,
		Limit:    req.Size,
	}
	ctx := context.Background()
	total, episodes, err := eh.operator.List(ctx, params)
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
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", episodes, pagination))
}

type searchEpsiodeRequest struct {
	Query string `query:"query"`
	Page  int    `query:"page"`
	Size  int    `query:"size"`
}

func (eh *EpisodeHandler) Search(c *fiber.Ctx) error {
	seasonId, err := c.ParamsInt("seasonId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &searchEpsiodeRequest{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Page = 10
	}
	params := operator.SearchEpisodeParams{
		SeasonID: seasonId,
		Query:    req.Query,
		Offset:   (req.Page - 1) * req.Size,
		Limit:    req.Size,
	}
	ctx := context.Background()
	total, episodes, err := eh.operator.Search(ctx, params)
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
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", episodes, pagination))
}

func (eh *EpisodeHandler) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	ctx := context.Background()
	episode, err := eh.operator.Get(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", episode, nil))
}

type updateEpisodeRequest struct {
	Name        string `json:"name"`
	Value       int64  `json:"value"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (eh *EpisodeHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse("Invalid params", err.Error()))
	}
	req := &updateEpisodeRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	params := operator.UpdateEpisodeParams{
		ID:          id,
		Name:        req.Name,
		Value:       req.Value,
		Description: req.Description,
		Status:      req.Status,
	}
	ctx := context.Background()
	episode, err := eh.operator.Update(ctx, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", episode, nil))
}

func (eh *EpisodeHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid params", err.Error())
	}
	ctx := context.Background()
	err = eh.operator.Delete(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Internal server error", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse("Success", nil, nil))
}
