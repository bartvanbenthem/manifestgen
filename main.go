package main

import (
	"flag"
	"fmt"

	"github.com/bartvanbenthem/manifestgen/app"
)

func main() {

	// init argument variables
	var value, template, output, filetype *string
	// set and parse flags
	value = flag.String("value", "./value.yaml", "path/file to values file")
	template = flag.String("template", "./template.yaml", "path/file to template file")
	output = flag.String("output", "./output.yaml", "path/file to output file")
	filetype = flag.String("filetype", "yaml", "yaml / json")
	flag.Parse()

	// check if file-type is yaml or json and run corresponding fenerate function
	if string(*filetype) == string("yaml") {
		var a app.ManifestGenClient
		a.GenerateYamlManifest(string(*value), string(*template), string(*output))
	} else {
		var a app.ManifestGenClient
		a.GenerateJSONManifest(string(*value), string(*template), string(*output))
	}
	// response
	fmt.Printf("%v manifest is generated\n", *output)
}
