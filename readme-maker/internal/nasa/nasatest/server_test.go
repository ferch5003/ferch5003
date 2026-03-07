package nasatest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer_NewServer(t *testing.T) {
	// When
	server := NewServer()
	defer server.Close()

	// Then
	require.NotNil(t, server)
	require.NotEmpty(t, server.URL)
}

func TestServer_ServeHTTPSuccessful(t *testing.T) {
	// Given
	server := httptest.NewServer(&Server{
		StatusCode:       http.StatusOK,
		ResponseBody:     `{"title":"test","date":"2024-01-01"}`,
		ResponseBodyJSON: true,
	})
	defer server.Close()

	// When
	resp, err := http.Get(server.URL + "/planetary/apod")
	require.NoError(t, err)
	defer func() { _, _ = io.Copy(io.Discard, resp.Body) }()

	// Then
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestServer_ServeHTTPError(t *testing.T) {
	// Given
	server := httptest.NewServer(&Server{
		StatusCode:       http.StatusInternalServerError,
		ResponseBody:     `{"error":"internal error"}`,
		ResponseBodyJSON: true,
	})
	defer server.Close()

	// When
	resp, err := http.Get(server.URL + "/planetary/apod")
	require.NoError(t, err)
	defer func() { _, _ = io.Copy(io.Discard, resp.Body) }()

	// Then
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestServer_ServeHTTPNotFound(t *testing.T) {
	// Given
	server := httptest.NewServer(&Server{
		StatusCode:       http.StatusOK,
		ResponseBody:     `{"title":"test"}`,
		ResponseBodyJSON: true,
	})
	defer server.Close()

	// When - request to a different endpoint
	resp, err := http.Get(server.URL + "/other/endpoint")
	require.NoError(t, err)
	defer func() { _, _ = io.Copy(io.Discard, resp.Body) }()

	// Then - should not write anything
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
