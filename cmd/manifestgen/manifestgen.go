package main

import (
	"flag"
	"log"

	"github.com/bartvanbenthem/manifestgen/pkg/manifest"
)

// init argument variables for manifestgen
var valuepath, template, output, filetype *string

// generate manifest function, takes an Generator interface as input
// argument variables are used to read, parse and write manifests
func generateManifest(g manifest.Builder) error {
	values, err := g.Parse(string(*valuepath))
	if err != nil {
		return err
	}
	err = g.Write(values, string(*template), string(*output))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var err error
	// set and parse flags
	filetype = flag.String("filetype", "json", "choose input type json / yaml")
	valuepath = flag.String("value", "", "path/file to values file")
	template = flag.String("template", "", "path/file to template file")
	output = flag.String("output", "./output.json", "path/file for writing output")
	flag.Parse()

	// check if file-type is yaml or json and run corresponding function
	if string(*filetype) == string("yaml") {
		err = generateManifest(&manifest.YAML{})
	} else if string(*filetype) == string("json") {
		err = generateManifest(&manifest.JSON{})
	}
	// response
	if err != nil {
		log.Printf("Error during %v manifest generation: %s\n", *output, err)
	} else {
		log.Printf("%v manifest is generated\n", *output)
	}
}
