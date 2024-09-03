package api

import (
	"github.com/dongwlin/anime-list/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type createAnimeRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

func (s *Server) createAnime(ctx *gin.Context) {
	req := &createAnimeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	arg := db.CreateAnimeParams{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	anime, err := s.store.CreateAnime(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create anime", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", anime, nil)
}

type listAnimeRequest struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

func (s *Server) listAnime(ctx *gin.Context) {
	req := &listAnimeRequest{}
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

	arg := db.ListAnimeParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}
	animes, err := s.store.ListAnime(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list anime", err.Error())
		return
	}
	total, err := s.store.CountAnime(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list anime", err.Error())
		return
	}
	totalPages := total / req.Size
	if total%req.Size != 0 {
		totalPages++
	}
	pagination := &Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		PerPage:     req.Size,
		CurrentPage: req.Page,
	}

	SuccessResponse(ctx, http.StatusOK, "Success", animes, pagination)
}

func (s *Server) getAnime(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	anime, err := s.store.GetAnime(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get anime", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", anime, nil)
}

type updateAnimeRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status int64  `json:"status"`
}

func (s *Server) updateAnime(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	req := &updateAnimeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	arg := db.UpdateAnimeParams{
		ID:          id,
		Name:        req.Name,
		Description: req.Desc,
		Status:      req.Status,
	}
	anime, err := s.store.UpdateAnime(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update anime", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", anime, nil)
}

func (s *Server) deleteAnime(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	err = s.store.DeleteAnime(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete anime", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", nil, nil)
}
