package nasa

import (
	"strings"
	"text/template"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
)

type nasaTemplate struct {
	client    Client
	Templates []templates.Templater
}

func NewNasaTemplate(client Client) templates.Templater {
	return &nasaTemplate{
		client: client,
	}
}

func (nt *nasaTemplate) AddTemplates(templates []templates.Templater) {
	nt.Templates = append(nt.Templates, templates...)
}

func (nt *nasaTemplate) Parse(in string) (string, error) {
	parsed := new(strings.Builder)

	apodParams := dto.APODRequestParams{}
	response, err := nt.client.GetAPOD(apodParams)
	if err != nil {
		return "", err
	}

	var nasaAPODValues struct {
		Nasa struct {
			APOD dto.APODResponse
		}
	}
	nasaAPODValues.Nasa.APOD = response

	tmpl, err := template.New("nasa").Parse(in)
	if err != nil {
		return "", err
	}

	if err = tmpl.Execute(parsed, nasaAPODValues); err != nil {
		return "", err
	}

	return parsed.String(), nil
}
