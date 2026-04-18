package nasa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/google/go-querystring/query"
)

const _planetaryPath = "planetary"

// Config holds the NASA API configuration.
type Config struct {
	BaseURL      string
	APIKey       string
	MaxRetries   int           // optional, defaults to 3
	RetryBackoff time.Duration // optional, defaults to 1s; doubles each attempt
}

type Client interface {
	// GetAPOD returns an APODResponse that the response for the /planetary/apod endpoint for NASA API.
	GetAPOD(params dto.APODRequestParams) (dto.APODResponse, error)
}

type client struct {
	baseUrl      string
	apiKey       string
	httpClient   *http.Client
	maxRetries   int
	retryBackoff time.Duration
}

// NewClient creates a new NASA API client with the given configuration.
func NewClient(cfg Config) Client {
	maxRetries := cfg.MaxRetries
	if maxRetries <= 0 {
		maxRetries = 3
	}
	backoff := cfg.RetryBackoff
	if backoff <= 0 {
		backoff = time.Second
	}
	return client{
		baseUrl: cfg.BaseURL,
		apiKey:  cfg.APIKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		maxRetries:   maxRetries,
		retryBackoff: backoff,
	}
}

func (c client) getWithRetry(url string) (*http.Response, error) {
	var lastErr error
	backoff := c.retryBackoff
	for attempt := 1; attempt <= c.maxRetries; attempt++ {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(backoff)
				backoff *= 2
			}
			continue
		}
		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			body, readErr := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if readErr != nil {
				body = []byte("(unable to read body)")
			}
			lastErr = fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
			if attempt < c.maxRetries {
				time.Sleep(backoff)
				backoff *= 2
			}
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("after %d attempts: %w", c.maxRetries, lastErr)
}

func (c client) GetAPOD(params dto.APODRequestParams) (dto.APODResponse, error) {
	// Transform the values of the APODRequestParams to a query.Values in order to facilitate the pass of query params
	// for the endpoint.
	queryValues, err := query.Values(params)
	if err != nil {
		return dto.APODResponse{}, err
	}

	apodEndpoint := fmt.Sprintf("%s/%s/%s?api_key=%s&%v", c.baseUrl, _planetaryPath, "apod", c.apiKey, queryValues.Encode())
	resp, err := c.getWithRetry(apodEndpoint)
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
