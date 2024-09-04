package api

import (
	"github.com/dongwlin/anime-list/internal/db"
	"github.com/dongwlin/anime-list/internal/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type createEpisodeRequest struct {
	Name        string `json:"name" binding:"required"`
	Value       int64  `json:"value" binding:"required"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

func (s *Server) createEpisode(ctx *gin.Context) {
	seasonId, err := strconv.ParseInt(ctx.Param("seasonId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	req := &createEpisodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	if !status.Valid(req.Status) {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", "Invalid status")
		return
	}

	arg := db.CreateEpisodeParams{
		SeasonID:    seasonId,
		Name:        req.Name,
		Value:       req.Value,
		Description: req.Description,
		Status:      req.Status,
	}
	episode, err := s.store.CreateEpisode(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create episode", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", episode, nil)
}

type listEpisodesRequest struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

func (s *Server) listEpisodes(ctx *gin.Context) {
	seasonId, err := strconv.ParseInt(ctx.Param("seasonId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	req := &listEpisodesRequest{}
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

	arg := db.ListEpisodeParams{
		SeasonID: seasonId,
		Limit:    req.Size,
		Offset:   (req.Page - 1) * req.Size,
	}
	episodes, err := s.store.ListEpisode(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list episodes", err.Error())
		return
	}
	total, err := s.store.CountEpisode(ctx, seasonId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list episodes", err.Error())
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

	SuccessResponse(ctx, http.StatusOK, "Success", episodes, pagination)
}

func (s *Server) getEpisode(ctx *gin.Context) {
	episodeId, err := strconv.ParseInt(ctx.Param("episodeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	episode, err := s.store.GetEpisode(ctx, episodeId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get episode", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", episode, nil)
}

type updateEpisodeRequest struct {
	Name        string `json:"name" binding:"required"`
	Value       int64  `json:"value" binding:"required"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

func (s *Server) updateEpisode(ctx *gin.Context) {
	episodeId, err := strconv.ParseInt(ctx.Param("episodeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	req := &updateEpisodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}
	if !status.Valid(req.Status) {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", "Invalid status")
		return
	}

	arg := db.UpdateEpisodeParams{
		ID:          episodeId,
		Name:        req.Name,
		Value:       req.Value,
		Description: req.Description,
		Status:      req.Status,
	}
	episode, err := s.store.UpdateEpisode(ctx, arg)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update episode", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", episode, nil)
}

func (s *Server) deleteEpisode(ctx *gin.Context) {
	episodeId, err := strconv.ParseInt(ctx.Param("episodeId"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "Invalid params", err.Error())
		return
	}

	err = s.store.DeleteEpisode(ctx, episodeId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete episode", err.Error())
		return
	}

	SuccessResponse(ctx, http.StatusOK, "Success", nil, nil)
}
