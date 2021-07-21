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

func (c *ManifestGenClient) ConvertJSON2Yaml(pathJSONFile string) {
}

func (c *ManifestGenClient) ConvertYaml2JSON(pathYamlFile string) {
}

func (c *ManifestGenClient) GenerateJSONManifest(pathValuesFile, pathTemplateFile, pathOutputFile string) {
	// open the values yaml file
	jsonFile, err := os.Open(pathValuesFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	defer jsonFile.Close()

	// read the values file and create a byte slice output
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal byte slice into the values map
	err = json.Unmarshal(byteValue, &values)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	genFile, err := os.Create(pathOutputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer genFile.Close()

	// Execute template var injection and write to file
	err = tpl.Execute(genFile, values)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

// Function to generate the configurable manifest from the values and template files
// pathValuesFile wants a string containing path to the values yaml file
// pathTemplateFile wants a string containing path to the template yaml file
// pathOutputFolder wants a string containing path to the output folder
func (c *ManifestGenClient) GenerateYamlManifest(pathValuesFile, pathTemplateFile, pathOutputFile string) {
	// open the values yaml file
	yamlFile, err := os.Open(pathValuesFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	defer yamlFile.Close()

	// read the values file and create a byte slice output
	byteValue, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal byte slice into the values map
	err = yaml.Unmarshal(byteValue, &values)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	genFile, err := os.Create(pathOutputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer genFile.Close()

	// Execute template var injection and write to file
	err = tpl.Execute(genFile, values)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

}
