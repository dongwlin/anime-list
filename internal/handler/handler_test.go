package handler

import (
	"testing"

	"github.com/dongwlin/anime-list/internal/db/mock"
	"go.uber.org/mock/gomock"
)

func createTestHandler(t *testing.T) *Handler {
	ctrl := gomock.NewController(t)
	mockStore := mock.NewMockStore(ctrl)
	handler := NewHandler(mockStore)
	return handler
}
