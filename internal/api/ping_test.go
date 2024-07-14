package api

import (
	"bytes"
	"github.com/dongwlin/anime-list/internal/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock.NewMockStore(ctrl)
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := "/ping"
	request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte{}))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	data, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)
	require.Equal(t, "pong!!!!!", string(data))
}
