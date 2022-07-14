package template

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Builder interface {
	ParseToStdout([]byte, string) error
	ParseToFile([]byte, string, string) error
	ReadFromFile(string) ([]byte, error)
}

type JSON struct{}

func (r *JSON) ParseToStdout(data []byte, templateFile string) error {
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

func (r *JSON) ParseToFile(data []byte, templatePath, file string) error {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}
	// unmarshal yaml byte slice into the values map
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	genFile, err := os.Create(file)
	if err != nil {
		log.Fatalln(err)
	}

	defer genFile.Close()

	// Execute template var injection and write to output file
	err = tpl.Execute(genFile, values)
	if err != nil {
		return err
	}

	return err
}

func (r *JSON) ReadFromFile(file string) ([]byte, error) {
	var values []byte
	// open the json/yaml file
	File, err := os.Open(file)
	if err != nil {
		return values, err
	}
	defer File.Close()
	// read the file and create a byte slice output
	values, err = ioutil.ReadAll(File)
	if err != nil {
		return values, err
	}

	return values, err
}

type YAML struct{}

func (r *YAML) ParseToStdout(data []byte, templateFile string) error {
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

func (r *YAML) ParseToFile(data []byte, templatePath, file string) error {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal yaml byte slice into the values map
	err := yaml.Unmarshal(data, &values)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	genFile, err := os.Create(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer genFile.Close()

	// Execute template var injection and write to output file
	err = tpl.Execute(genFile, values)
	if err != nil {
		return err
	}

	return err
}

func (r *YAML) ReadFromFile(file string) ([]byte, error) {
	var values []byte
	// open the json/yaml file
	File, err := os.Open(file)
	if err != nil {
		return values, err
	}
	defer File.Close()
	// read the file and create a byte slice output
	values, err = ioutil.ReadAll(File)
	if err != nil {
		return values, err
	}

	return values, err
}
