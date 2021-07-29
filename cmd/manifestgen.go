package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bartvanbenthem/manifestgen/core"
)

// init argument variables for manifestgen
var valuepath, template, output, filetype *string

// generate manifest function, takes an Generator interface as input
// argument variables are used to read, parse and write manifests
func generateManifest(g core.Generator) {
	values, err := g.Parse(string(*valuepath))
	if err != nil {
		log.Println(err)
	}
	err = g.Write(values, string(*template), string(*output))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// set and parse flags
	valuepath = flag.String("value", "./value.yaml", "path/file to values file")
	template = flag.String("template", "./template.yaml", "path/file to template file")
	output = flag.String("output", "./output.yaml", "path/file to output file")
	filetype = flag.String("filetype", "yaml", "yaml / json")
	flag.Parse()
	// check if file-type is yaml or json and run corresponding function
	if string(*filetype) == string("yaml") {
		generateManifest(&core.YAMLClient{})
	} else if string(*filetype) == string("json") {
		generateManifest(&core.JSONClient{})
	}
	// response
	fmt.Printf("%v manifest is generated\n", *output)
}
