package handler

import "github.com/dongwlin/anime-list/internal/db"

type Handler struct {
	store db.Store
}

func NewHandler(store db.Store) *Handler {
	return &Handler{
		store: store,
	}
}
