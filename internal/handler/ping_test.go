package handler

import (
	"github.com/dongwlin/anime-list/internal/store/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	app := fiber.New()
	ctrl := gomock.NewController(t)
	mockStore := mock.NewMockStore(ctrl)
	handler := NewHandler(mockStore)
	pingHandler := NewPingHandler(handler)
	app.All("/ping", pingHandler.Ping)
	methods := []string{"GET", "POST", "PUT", "DELETE"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/ping", nil)
			resp, err := app.Test(req)
			require.NoError(t, err)
			require.Equal(t, fiber.StatusOK, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, "pong!!!!!", string(body))
		})
	}
}
