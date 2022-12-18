package templates

// Templater is in charge to parse the data in a string with a new one.
type Templater interface {
	AddTemplates(templates []Templater)
	Parse(in string) (parsed string, err error)
}
