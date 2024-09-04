package api

import (
	"github.com/dongwlin/anime-list/internal/db"
	"github.com/dongwlin/anime-list/internal/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type createSeasonRequest struct {
	Name        string    `json:"name" binding:"required"`
	Value       int64     `json:"value" binding:"required"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int64     `json:"status"`
}

func (s *Server) createSeason(ctx *gin.Context) {
	animeId, err := strconv.ParseInt(ctx.Param("animeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	req := &createSeasonRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	if !status.Valid(req.Status) {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", "Invalid status")
		return
	}

	arg := db.CreateSeasonParams{
		AnimeID:     animeId,
		Name:        req.Name,
		Value:       req.Value,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	season, err := s.store.CreateSeason(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create season", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", season, nil)
}

type listSeasonsRequest struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

func (s *Server) listSeason(ctx *gin.Context) {
	animeId, err := strconv.ParseInt(ctx.Param("animeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	req := &listSeasonsRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	arg := db.ListSeasonParams{
		AnimeID: animeId,
		Limit:   req.Size,
		Offset:  (req.Page - 1) * req.Size,
	}
	seasons, err := s.store.ListSeason(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list seasons", err.Error())
		return
	}
	total, err := s.store.CountSeason(ctx, animeId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to count seasons", err.Error())
		return
	}
	totalPages := total / req.Size
	if total%req.Size != 0 {
		totalPages++
	}
	if totalPages == 0 {
		req.Page = 0
	}
	pagination := &Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		PerPage:     req.Size,
		CurrentPage: req.Page,
	}

	SuccessResponse(ctx, http.StatusOK, "Success", seasons, pagination)
}

func (s *Server) getSeason(ctx *gin.Context) {
	seasonId, err := strconv.ParseInt(ctx.Param("seasonId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	season, err := s.store.GetSeason(ctx, seasonId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get season", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", season, nil)
}

type updateSeasonRequest struct {
	Name        string    `json:"name" binding:"required"`
	Value       int64     `json:"value" binding:"required"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int64     `json:"status"`
}

func (s *Server) updateSeason(ctx *gin.Context) {
	seasonId, err := strconv.ParseInt(ctx.Param("seasonId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	req := &updateSeasonRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	if !status.Valid(req.Status) {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", "Invalid status")
		return
	}

	arg := db.UpdateSeasonParams{
		ID:          seasonId,
		Name:        req.Name,
		Value:       req.Value,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	season, err := s.store.UpdateSeason(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update season", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", season, nil)
}

func (s *Server) deleteSeason(ctx *gin.Context) {
	seasonId, err := strconv.ParseInt(ctx.Param("seasonId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	err = s.store.DeleteSeason(ctx, seasonId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete season", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", nil, nil)
}
