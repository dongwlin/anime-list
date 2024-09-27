package handler

import "github.com/dongwlin/anime-list/internal/store"

type Handler struct {
	store store.Store
}

func NewHandler(store store.Store) *Handler {
	return &Handler{
		store: store,
	}
}
