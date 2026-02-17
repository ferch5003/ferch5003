package nasa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/google/go-querystring/query"
)

const _planetaryPath = "planetary"

type Client interface {
	// GetAPOD returns an APODResponse that the response for the /planetary/apod endpoint for NASA API.
	GetAPOD(params dto.APODRequestParams) (dto.APODResponse, error)
}

type client struct {
	baseUrl    string
	apiKey     string
	httpClient *http.Client
}

func NewClient() Client {
	return client{
		baseUrl: os.Getenv("NASA_BASE_URL"),
		apiKey:  os.Getenv("NASA_API_KEY"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c client) GetAPOD(params dto.APODRequestParams) (dto.APODResponse, error) {
	// Transform the values of the APODRequestParams to a query.Values in order to facilitate the pass of query params
	// for the endpoint.
	queryValues, err := query.Values(params)
	if err != nil {
		return dto.APODResponse{}, err
	}

	apodEndpoint := fmt.Sprintf("%s/%s/%s?api_key=%s&%v", c.baseUrl, _planetaryPath, "apod", c.apiKey, queryValues.Encode())
	resp, err := c.httpClient.Get(apodEndpoint)
	if err != nil {
		return dto.APODResponse{}, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}()

	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return dto.APODResponse{}, fmt.Errorf("HTTP error %d: failed to read response body", resp.StatusCode)
		}
		return dto.APODResponse{}, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.APODResponse{}, err
	}

	var apodResponse dto.APODResponse
	err = json.Unmarshal(body, &apodResponse)
	if err != nil {
		return dto.APODResponse{}, err
	}

	return apodResponse, nil
}
