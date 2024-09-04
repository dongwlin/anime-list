package api

import (
	"github.com/dongwlin/anime-list/internal/db"
	"github.com/dongwlin/anime-list/internal/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type createTheaterRequest struct {
	Name        string    `json:"name" binding:"required"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int64     `json:"status"`
}

func (s *Server) createTheater(ctx *gin.Context) {
	animeId, err := strconv.ParseInt(ctx.Param("animeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}
	req := &createTheaterRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}
	if !status.Valid(req.Status) {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", "Invalid status")
		return
	}

	arg := db.CreateTheaterParams{
		AnimeID:     animeId,
		Name:        req.Name,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	theater, err := s.store.CreateTheater(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create theater", nil)
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", theater, nil)
}

type listTheaterRequest struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

func (s *Server) listTheater(ctx *gin.Context) {
	animeId, err := strconv.ParseInt(ctx.Param("animeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}
	req := &listTheaterRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	arg := db.ListTheaterParams{
		AnimeID: animeId,
		Limit:   req.Size,
		Offset:  (req.Page - 1) * req.Size,
	}
	theaters, err := s.store.ListTheater(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list theater", nil)
		return
	}
	total, err := s.store.CountTheater(ctx, animeId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list theater", nil)
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

	SuccessResponse(ctx, http.StatusOK, "Success", theaters, pagination)
}

func (s *Server) getTheater(ctx *gin.Context) {
	theaterId, err := strconv.ParseInt(ctx.Param("theaterId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}

	theater, err := s.store.GetTheater(ctx, theaterId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get theater", nil)
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", theater, nil)
}

type updateTheaterRequest struct {
	Name        string    `json:"name" binding:"required"`
	Cover       string    `json:"cover"`
	ReleasedAt  time.Time `json:"released_at"`
	Description string    `json:"description"`
	Status      int64     `json:"status"`
}

func (s *Server) updateTheater(ctx *gin.Context) {
	theaterId, err := strconv.ParseInt(ctx.Param("theaterId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}
	req := &updateTheaterRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}
	if !status.Valid(req.Status) {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", "Invalid status")
		return
	}

	arg := db.UpdateTheaterParams{
		ID:          theaterId,
		Name:        req.Name,
		Cover:       req.Cover,
		ReleasedAt:  req.ReleasedAt,
		Description: req.Description,
		Status:      req.Status,
	}
	theater, err := s.store.UpdateTheater(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update theater", nil)
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", theater, nil)
}

func (s *Server) deleteTheater(ctx *gin.Context) {
	theaterId, err := strconv.ParseInt(ctx.Param("theaterId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", nil)
		return
	}

	err = s.store.DeleteTheater(ctx, theaterId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete theater", nil)
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", nil, nil)
}
