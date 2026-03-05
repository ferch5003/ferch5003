package nasa

import (
	"testing"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/nasatest"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
	"github.com/stretchr/testify/require"
)

func TestNasaTemplate_ParseSuccessful(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	nasaClient := NewClient(Config{
		BaseURL: server.URL,
		APIKey:  "test",
	})
	nasaTemplate := NewNasaTemplate(nasaClient)

	template := "{{.Nasa.APOD.Title}} template for nasa"

	// When
	templateString, err := nasaTemplate.Parse(template)

	// Then
	require.NoError(t, err)
	require.Equal(t, "test template for nasa", templateString)
}

func TestNasaTemplate_ParseErrorWrongURL(t *testing.T) {
	// Given
	nasaClient := NewClient(Config{
		BaseURL: "notexist",
		APIKey:  "test",
	})
	nasaTemplate := NewNasaTemplate(nasaClient)

	template := "{{.Nasa.APOD.Title}} template for nasa"

	// When
	templateString, err := nasaTemplate.Parse(template)

	// Then
	require.ErrorContains(t, err, "unsupported protocol scheme")
	require.Equal(t, "", templateString)
}

func TestNasaTemplate_AddTemplates(t *testing.T) {
	nasaTemplate := NewNasaTemplate(nil)

	mockTemplate := &mockTemplater{}
	nasaTemplate.AddTemplates([]templates.Templater{mockTemplate})
}

func TestNasaAPODValues_IsVideoFormatWithYoutube(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://youtube.com/watch?v=123",
	}

	result := values.IsVideoFormat()

	require.True(t, result)
}

func TestNasaAPODValues_IsVideoFormatWithoutYoutube(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://example.com/image.jpg",
	}

	result := values.IsVideoFormat()

	require.False(t, result)
}

func TestNasaTemplate_ParseErrorInvalidTemplate(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	nasaClient := NewClient(Config{
		BaseURL: server.URL,
		APIKey:  "test",
	})
	nasaTemplate := NewNasaTemplate(nasaClient)

	template := "{{.Invalid"

	// When
	templateString, err := nasaTemplate.Parse(template)

	// Then
	require.Error(t, err)
	require.Equal(t, "", templateString)
}

func TestNasaTemplate_ParseErrorClientError(t *testing.T) {
	// Given
	nasaClient := NewClient(Config{
		BaseURL: "notexist",
		APIKey:  "test",
	})
	nasaTemplate := NewNasaTemplate(nasaClient)

	template := "{{.Nasa.APOD.Title}}"

	// When
	templateString, err := nasaTemplate.Parse(template)

	// Then
	require.Error(t, err)
	require.Equal(t, "", templateString)
}

type mockTemplater struct{}

func (m *mockTemplater) AddTemplates(templates []templates.Templater) {}
func (m *mockTemplater) Parse(in string) (string, error)              { return in, nil }

func TestNasaAPODValues_GetYouTubeIDWithWatchURL(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}

	result := values.GetYouTubeID()

	require.Equal(t, "dQw4w9WgXcQ", result)
}

func TestNasaAPODValues_GetYouTubeIDWithShortURL(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://youtu.be/dQw4w9WgXcQ",
	}

	result := values.GetYouTubeID()

	require.Equal(t, "dQw4w9WgXcQ", result)
}

func TestNasaAPODValues_GetYouTubeIDWithEmbedURL(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://www.youtube.com/embed/dQw4w9WgXcQ",
	}

	result := values.GetYouTubeID()

	require.Equal(t, "dQw4w9WgXcQ", result)
}

func TestNasaAPODValues_GetYouTubeIDWithNonYouTubeURL(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://example.com/video.mp4",
	}

	result := values.GetYouTubeID()

	require.Equal(t, "", result)
}

func TestNasaAPODValues_GetYouTubeIDWithEmptyURL(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "",
	}

	result := values.GetYouTubeID()

	require.Equal(t, "", result)
}

func TestNasaAPODValues_IsVideoFormatWithMediaTypeVideo(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		MediaType: "video",
		Url:       "https://apod.nasa.gov/apod/image/2603/video.mp4",
	}

	result := values.IsVideoFormat()

	require.True(t, result)
}

func TestNasaAPODValues_IsVideoFormatWithMp4Url(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		MediaType: "image",
		Url:       "https://apod.nasa.gov/apod/image/2603/FlyingNorth_MarsExpress.mp4",
	}

	result := values.IsVideoFormat()

	require.True(t, result)
}

func TestNasaAPODValues_IsVideoFormatWithMovUrl(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		MediaType: "image",
		Url:       "https://example.com/video.mov",
	}

	result := values.IsVideoFormat()

	require.True(t, result)
}

func TestNasaAPODValues_IsVideoFormatWithWebMUrl(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		MediaType: "image",
		Url:       "https://example.com/video.webm",
	}

	result := values.IsVideoFormat()

	require.True(t, result)
}

func TestNasaAPODValues_IsYouTubeVideo(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://youtube.com/watch?v=123",
	}

	result := values.IsYouTubeVideo()

	require.True(t, result)
}

func TestNasaAPODValues_IsYouTubeVideoFalse(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "https://apod.nasa.gov/apod/video.mp4",
	}

	result := values.IsYouTubeVideo()

	require.False(t, result)
}

func TestNasaAPODValues_IsYouTubeVideoEmptyUrl(t *testing.T) {
	var values _nasaAPODValues
	values.Nasa.APOD = dto.APODResponse{
		Url: "",
	}

	result := values.IsYouTubeVideo()

	require.False(t, result)
}
