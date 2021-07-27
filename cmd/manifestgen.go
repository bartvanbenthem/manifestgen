package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bartvanbenthem/manifestgen/app"
)

// init generator interface
var generator app.Generator

// init argument variables
var valuepath, template, output, filetype *string

func main() {
	// set and parse flags
	valuepath = flag.String("value", "./value.yaml", "path/file to values file")
	template = flag.String("template", "./template.yaml", "path/file to template file")
	output = flag.String("output", "./output.yaml", "path/file to output file")
	filetype = flag.String("filetype", "yaml", "yaml / json")
	flag.Parse()

	// check if file-type is yaml or json and run corresponding function
	if string(*filetype) == string("yaml") {
		generator = app.YAMLClient{}
		values, err := generator.Parser(string(*valuepath))
		if err != nil {
			log.Println(err)
		}
		err = generator.Writer(values, string(*template), string(*output))
		if err != nil {
			log.Println(err)
		}
	} else {
		generator = app.JSONClient{}
		values, err := generator.Parser(string(*valuepath))
		if err != nil {
			log.Println(err)
		}
		err = generator.Writer(values, string(*template), string(*output))
		if err != nil {
			log.Println(err)
		}
	}
	// response
	fmt.Printf("%v manifest is generated\n", *output)
}
