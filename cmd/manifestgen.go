package main

import (
	"flag"
	"fmt"

	"github.com/bartvanbenthem/manifestgen/pkg"
)

// init argument variables for manifestgen
var valuepath, template, output, filetype *string

// generate manifest function, takes an Generator interface as input
// argument variables are used to read, parse and write manifests
func generateManifest(g pkg.Generator) error {
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
	valuepath = flag.String("value", "./value.yaml", "path/file to values file")
	template = flag.String("template", "./template.yaml", "path/file to template file")
	output = flag.String("output", "./output.yaml", "path/file to output file")
	filetype = flag.String("filetype", "yaml", "yaml / json")
	flag.Parse()

	// check if file-type is yaml or json and run corresponding function
	if string(*filetype) == string("yaml") {
		err = generateManifest(&pkg.YAMLClient{})
	} else if string(*filetype) == string("json") {
		err = generateManifest(&pkg.JSONClient{})
	}
	// response
	if err != nil {
		fmt.Printf("Error during %v manifest generation: %s\n", *output, err)
	} else {
		fmt.Printf("%v manifest is generated\n", *output)
	}
}
