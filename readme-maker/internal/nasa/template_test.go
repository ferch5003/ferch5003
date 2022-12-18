package nasa

import (
	"testing"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/nasatest"
	"github.com/stretchr/testify/require"
)

func TestNasaTemplate_ParseSuccessful(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	nasaClient := NewClient()
	nasaTemplate := NewNasaTemplate(nasaClient)

	template := "{{.Title}} template for nasa"

	// When
	templateString, err := nasaTemplate.Parse(template)

	// Then
	require.NoError(t, err)
	require.Equal(t, "test template for nasa", templateString)
}

func TestNasaTemplate_ParseErrorWrongURL(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	t.Setenv("NASA_BASE_URL", "notexist")

	nasaClient := NewClient()
	nasaTemplate := NewNasaTemplate(nasaClient)

	template := "{{.Title}} template for nasa"

	// When
	templateString, err := nasaTemplate.Parse(template)

	// Then
	require.ErrorContains(t, err, "unsupported protocol scheme")
	require.Equal(t, "", templateString)
}
