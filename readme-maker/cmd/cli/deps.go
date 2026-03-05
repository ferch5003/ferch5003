package main

import (
	"os"
	"path/filepath"

	nasa2 "github.com/ferch5003/ferch5003/readme-maker/internal/nasa"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/storage"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
	"github.com/joho/godotenv"
)

type dependencies struct {
	Storage   storage.Storage
	Templates []templates.Templater
}

func newDependencies() dependencies {
	// Load .env file if it exists (useful for local development)
	// Will not fail if file doesn't exist
	envPath := filepath.Join("readme-maker", ".env")
	_ = godotenv.Load(envPath)

	fs := storage.New("./README.md.tpl", "./README.md")

	// Slice to save all the templates.
	var templates []templates.Templater

	nasaClient := nasa2.NewClient(nasa2.Config{
		BaseURL: os.Getenv("NASA_BASE_URL"),
		APIKey:  os.Getenv("NASA_API_KEY"),
	})
	nasaTemplate := nasa2.NewNasaTemplate(nasaClient)

	// Save dependency templates.
	templates = append(templates, nasaTemplate)

	deps := dependencies{
		Storage:   fs,
		Templates: templates,
	}

	return deps
}
