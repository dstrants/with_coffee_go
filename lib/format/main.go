package format

import (
	"text/template"
)

func LoadTemplate(name string, paths []string) (*template.Template, error) {

	temp, err := template.New(name).ParseFiles(paths...)
	if err != nil {
		return temp, err
	}
	return temp, nil
}
