package app

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ManifestGenClient struct {
}

type Generator interface {
	GenerateManifest(pathValuesFile, pathTemplateFile, pathOutputFile string)
}

// Function to generate manifest from the values and template files
// pathValuesFile wants a string containing path and file name to the values yaml file
// pathTemplateFile wants a string containing path and file name to the template yaml file
// pathOutputFile wants a string containing path and file name to the yaml output file
func (c *ManifestGenClient) GenerateYamlManifest(pathValuesFile, pathTemplateFile, pathOutputFile string) error {
	// open the yaml file
	yamlFile, err := os.Open(pathValuesFile)
	if err != nil {
		return err
	}
	defer yamlFile.Close()
	// read the file and create a byte slice output
	byteValue, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return err
	}
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}
	// unmarshal byte slice into the values map
	err = yaml.Unmarshal(byteValue, &values)
	if err != nil {
		return err
	}
	//
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

// Function to generate manifest from the values and template files
// pathValuesFile wants a string containing path and file name to the values json file
// pathTemplateFile wants a string containing path and file name to the template json file
// pathOutputFile wants a string containing path and file name to the json output file
func (c *ManifestGenClient) GenerateJSONManifest(pathValuesFile, pathTemplateFile, pathOutputFile string) error {
	// open the json file
	jsonFile, err := os.Open(pathValuesFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	// read the file and create a byte slice output
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}
	// unmarshal byte slice into the values map
	err = json.Unmarshal(byteValue, &values)
	if err != nil {
		return err
	}
	//
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
