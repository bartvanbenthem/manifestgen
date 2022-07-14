package template

import (
	"encoding/json"
	"html/template"
	"os"

	"gopkg.in/yaml.v2"
)

type Builder interface {
	TemplateParser(data []byte, templateFile string) error
	ReadFromFile(file string) ([]byte, error)
	WriteToFile(data []byte, file string) error
}

type JSON struct{}

func (r *JSON) TemplateParser(data []byte, templateFile string) error {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal yaml byte slice into the values map
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(templateFile)
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

func (r *JSON) ReadFromFile(file string) ([]byte, error) {
	var output []byte
	var err error
	return output, err
}

func (r *JSON) WriteToFile(data []byte, file string) error {
	var err error
	return err
}

type YAML struct{}

func (r *YAML) TemplateParser(data []byte, templateFile string) error {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal yaml byte slice into the values map
	err := yaml.Unmarshal(data, &values)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(templateFile)
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

func (r *YAML) ReadFromFile(file string) ([]byte, error) {
	var output []byte
	var err error
	return output, err
}

func (r *YAML) WriteToFile(data []byte, file string) error {
	var err error
	return err
}
