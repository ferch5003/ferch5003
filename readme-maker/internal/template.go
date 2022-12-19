package internal

import (
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
)

type template struct {
	templateString string
	Templates      []templates.Templater
}

func NewTemplate() templates.Templater {
	return &template{}
}

func (t *template) AddTemplates(templates []templates.Templater) {
	t.Templates = append(t.Templates, templates...)
}

func (t *template) Parse(in string) (string, error) {
	// templateString receive the incoming template (the raw one) in order to parse the template.
	t.templateString = in

	var err error
	for _, template := range t.Templates {
		t.templateString, err = template.Parse(t.templateString)
		if err != nil {
			return "", err
		}
	}
	return t.templateString, nil
}
