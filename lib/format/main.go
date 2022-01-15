package format

import (
	"text/template"
)

// Loads a template based on tha name and paths
func LoadTemplate(name string, paths []string) (*template.Template, error) {

	temp, err := template.New(name).ParseFiles(paths...)
	if err != nil {
		return temp, err
	}
	return temp, nil
}
