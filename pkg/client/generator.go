package client

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Generator interface {
	Parse(pathValuesFile string) (map[string]interface{}, error)
	Write(values map[string]interface{}, pathTemplateFile, pathOutputFile string) error
}

type YAML struct {
}

// Function to generate manifest from the values and template files
// pathValuesFile wants a string containing path and file name to the values yaml file
// pathTemplateFile wants a string containing path and file name to the template yaml file
// pathOutputFile wants a string containing path and file name to the yaml output file
func (c *YAML) Parse(pathValuesFile string) (map[string]interface{}, error) {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}
	// open the json file
	yamlFile, err := os.Open(pathValuesFile)
	if err != nil {
		return values, err
	}
	defer yamlFile.Close()
	// read the file and create a byte slice output
	byteValue, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return values, err
	}
	// unmarshal byte slice into the values map
	err = yaml.Unmarshal(byteValue, &values)
	if err != nil {
		return values, err
	}
	return values, err
}

func (c *YAML) Write(values map[string]interface{}, pathTemplateFile, pathOutputFile string) error {
	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		return err
	}
	//
	genFile, err := os.Create(pathOutputFile)
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

type JSON struct {
}

// Function to generate manifest from the values and template files
// pathValuesFile wants a string containing path and file name to the values json file
// pathTemplateFile wants a string containing path and file name to the template json file
// pathOutputFile wants a string containing path and file name to the json output file
func (c *JSON) Parse(pathValuesFile string) (map[string]interface{}, error) {
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}
	// open the json file
	jsonFile, err := os.Open(pathValuesFile)
	if err != nil {
		return values, err
	}
	defer jsonFile.Close()
	// read the file and create a byte slice output
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return values, err
	}
	// unmarshal byte slice into the values map
	err = json.Unmarshal(byteValue, &values)
	if err != nil {
		return values, err
	}
	return values, err
}

func (c *JSON) Write(values map[string]interface{}, pathTemplateFile, pathOutputFile string) error {
	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		return err
	}
	//
	genFile, err := os.Create(pathOutputFile)
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
