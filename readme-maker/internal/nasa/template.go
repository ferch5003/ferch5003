package nasa

import (
	"regexp"
	"strings"
	"text/template"

	"github.com/ferch5003/ferch5003/readme-maker/internal/nasa/dto"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/templates"
)

var _videoFormats = []string{"youtube"}

var _directVideoExtensions = []string{".mp4", ".mov", ".webm"}

type _nasaAPODValues struct {
	Nasa struct {
		APOD dto.APODResponse
	}
}

func (n _nasaAPODValues) IsVideoFormat() bool {
	// Check media_type from NASA API first (most reliable)
	if n.Nasa.APOD.MediaType == "video" {
		return true
	}

	// Fallback: check URL for known video patterns
	for _, format := range _videoFormats {
		if n.Nasa.APOD.Url != "" && strings.Contains(n.Nasa.APOD.Url, format) {
			return true
		}
	}

	// Check for direct video file extensions
	for _, ext := range _directVideoExtensions {
		if n.Nasa.APOD.Url != "" && strings.HasSuffix(strings.ToLower(n.Nasa.APOD.Url), ext) {
			return true
		}
	}

	return false
}

func (n _nasaAPODValues) IsYouTubeVideo() bool {
	if n.Nasa.APOD.Url == "" {
		return false
	}
	return strings.Contains(n.Nasa.APOD.Url, "youtube")
}

// GetCopyright returns the copyright string cleaned for safe HTML usage.
// It replaces newlines and multiple spaces with single spaces.
func (n _nasaAPODValues) GetCopyright() string {
	copyright := n.Nasa.APOD.Copyright
	if copyright == "" {
		return ""
	}
	// Replace newlines and tabs with spaces
	copyright = strings.ReplaceAll(copyright, "\n", " ")
	copyright = strings.ReplaceAll(copyright, "\r", " ")
	copyright = strings.ReplaceAll(copyright, "\t", " ")
	// Replace multiple spaces with single space
	re := regexp.MustCompile(`\s+`)
	copyright = re.ReplaceAllString(copyright, " ")
	// Trim leading/trailing spaces
	return strings.TrimSpace(copyright)
}

func (n _nasaAPODValues) GetYouTubeID() string {
	url := n.Nasa.APOD.Url
	if url == "" {
		return ""
	}

	// Pattern for youtube.com/watch?v=VIDEO_ID
	re := regexp.MustCompile(`(?:youtube\.com/watch\?v=|youtu\.be/)([a-zA-Z0-9_-]{11})`)
	matches := re.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}

	// Pattern for youtube.com/embed/VIDEO_ID
	re = regexp.MustCompile(`youtube\.com/embed/([a-zA-Z0-9_-]{11})`)
	matches = re.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

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

	var nasaAPODValues _nasaAPODValues
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
