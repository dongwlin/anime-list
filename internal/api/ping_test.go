package api

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	server := NewServer()
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
