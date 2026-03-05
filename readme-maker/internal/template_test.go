package internal

import (
	"testing"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa"
	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/nasatest"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
	"github.com/stretchr/testify/require"
)

func TestTemplate_ParseSuccessful(t *testing.T) {
	server := nasatest.NewServer()
	defer server.Close()

	// Given
	nasaClient := nasa.NewClient(nasa.Config{
		BaseURL: server.URL,
		APIKey:  "test",
	})
	nasaTemplate := nasa.NewNasaTemplate(nasaClient)

	mainTemplate := NewTemplate()
	mainTemplate.AddTemplates([]templates.Templater{nasaTemplate})

	template := "{{.Nasa.APOD.Title}} template from main template"

	// When
	templateString, err := mainTemplate.Parse(template)

	// Then
	require.NoError(t, err)
	require.Equal(t, "test template from main template", templateString)
}

func TestTemplate_ParseErrorWrongURL(t *testing.T) {
	// Given
	nasaClient := nasa.NewClient(nasa.Config{
		BaseURL: "notexist",
		APIKey:  "test",
	})
	nasaTemplate := nasa.NewNasaTemplate(nasaClient)

	mainTemplate := NewTemplate()
	mainTemplate.AddTemplates([]templates.Templater{nasaTemplate})

	template := "{{.Nasa.APOD.Title}} template from main template"

	// When
	templateString, err := mainTemplate.Parse(template)

	// Then
	require.ErrorContains(t, err, "unsupported protocol scheme")
	require.Equal(t, "", templateString)
}
