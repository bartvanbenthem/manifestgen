package convert

import (
	"encoding/json"
	"html/template"
	"os"

	"gopkg.in/yaml.v2"
)

type Builder interface {
	TemplateParser(input []byte, pathTemplateFile string) error
}

type JSON struct{}

func (r *JSON) TemplateParser(input []byte, pathTemplateFile string) error {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal yaml byte slice into the values map
	err := json.Unmarshal(input, &values)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		return err
	}

	w := os.Stdout

	// Execute template var injection and write to output file
	err = tpl.Execute(w, values)
	if err != nil {
		return err
	}
	return err
}

type YAML struct{}

func (r *YAML) TemplateParser(input []byte, pathTemplateFile string) error {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal yaml byte slice into the values map
	err := yaml.Unmarshal(input, &values)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		return err
	}

	w := os.Stdout

	// Execute template var injection and write to output file
	err = tpl.Execute(w, values)
	if err != nil {
		return err
	}
	return err
}
