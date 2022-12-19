package main

import (
	nasa2 "github.com/ferch5003/ferch5003/readme-maker/internal/nasa"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/storage"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
)

type dependencies struct {
	Storage   storage.Storage
	Templates []templates.Templater
}

func newDependencies() (dependencies, error) {
	fs := storage.New("./README.md.tpl", "./README.md")

	// Slice to save all the templates.
	var templates []templates.Templater

	nasaClient := nasa2.NewClient()
	nasaTemplate := nasa2.NewNasaTemplate(nasaClient)

	// Save dependency templates.
	templates = append(templates, nasaTemplate)

	deps := dependencies{
		Storage:   fs,
		Templates: templates,
	}

	return deps, nil
}
