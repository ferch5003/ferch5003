package nasa

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/nasatest"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAPODSuccessfulRequest(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient(Config{
		BaseURL: server.URL,
		APIKey:  "test",
	})

	// When
	response, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.NoError(t, err)
	require.Equal(t, "test", response.Copyright)
	require.Equal(t, "2006-01-01", response.Date)
	require.Equal(t, "test", response.Explanation)
	require.Equal(t, "test", response.Hdurl)
	require.Equal(t, "test", response.MediaType)
	require.Equal(t, "test", response.ServiceVersion)
	require.Equal(t, "test", response.Title)
	require.Equal(t, "test", response.Url)
}

func TestClient_GetAPODErrorWrongURL(t *testing.T) {
	// Given
	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient(Config{
		BaseURL: "notexist",
		APIKey:  "test",
	})

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.ErrorContains(t, err, "unsupported protocol scheme")
}

func TestClient_GetAPODErrorInvalidJSON(t *testing.T) {
	server := httptest.NewServer(&nasatest.Server{
		StatusCode:       http.StatusOK,
		ResponseBody:     "not valid json",
		ResponseBodyJSON: false,
	})
	defer server.Close()

	// Given
	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient(Config{
		BaseURL: server.URL,
		APIKey:  "test",
	})

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.ErrorContains(t, err, "invalid character")
}

func TestClient_GetAPODErrorServerError(t *testing.T) {
	// Given
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer server.Close()

	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient(Config{
		BaseURL:      server.URL,
		APIKey:       "test",
		MaxRetries:   2,
		RetryBackoff: time.Millisecond,
	})

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.Error(t, err)
	require.Equal(t, 2, calls)
	require.ErrorContains(t, err, "HTTP error 500")
}

func TestClient_GetAPODErrorClientError(t *testing.T) {
	server := httptest.NewServer(&nasatest.Server{
		StatusCode:       http.StatusBadRequest,
		ResponseBody:     `{"error": "bad request"}`,
		ResponseBodyJSON: true,
	})
	defer server.Close()

	// Given
	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient(Config{
		BaseURL: server.URL,
		APIKey:  "test",
	})

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.ErrorContains(t, err, "HTTP error 400")
}

func TestClient_GetAPODFailsAfterRetriesExhausted(t *testing.T) {
	// Given
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("upstream connect error"))
	}))
	defer server.Close()

	nasaClient := NewClient(Config{
		BaseURL:      server.URL,
		APIKey:       "test",
		MaxRetries:   3,
		RetryBackoff: time.Millisecond,
	})

	// When
	_, err := nasaClient.GetAPOD(dto.APODRequestParams{})

	// Then
	require.Error(t, err)
	require.Equal(t, 3, calls)
	require.ErrorContains(t, err, "after 3 attempts")
	require.ErrorContains(t, err, "HTTP error 503")
}

func TestClient_GetAPODRetriesOn5xxThenSucceeds(t *testing.T) {
	// Given
	var calls int
	successBody := `{
		"copyright":"test",
		"date":"2006-01-01",
		"explanation":"test",
		"hdurl":"test",
		"media_type":"test",
		"service_version":"test",
		"title":"test",
		"url":"test"
	}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("upstream connect error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(successBody))
	}))
	defer server.Close()

	nasaClient := NewClient(Config{
		BaseURL:      server.URL,
		APIKey:       "test",
		MaxRetries:   3,
		RetryBackoff: time.Millisecond,
	})

	// When
	response, err := nasaClient.GetAPOD(dto.APODRequestParams{})

	// Then
	require.NoError(t, err)
	require.Equal(t, 3, calls)
	require.Equal(t, "test", response.Title)
}
