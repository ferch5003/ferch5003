package nasa

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/nasatest"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAPODSuccessfulRequest(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient()

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
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	t.Setenv("NASA_BASE_URL", "notexist")

	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient()

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
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient()

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.ErrorContains(t, err, "invalid character")
}

func TestClient_GetAPODErrorServerError(t *testing.T) {
	server := httptest.NewServer(&nasatest.Server{
		StatusCode:       http.StatusInternalServerError,
		ResponseBody:     `{"error": "internal server error"}`,
		ResponseBodyJSON: true,
	})
	defer server.Close()

	// Given
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient()

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
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
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	apodParams := dto.APODRequestParams{}
	nasaClient := NewClient()

	// When
	_, err := nasaClient.GetAPOD(apodParams)

	// Then
	require.ErrorContains(t, err, "HTTP error 400")
}
