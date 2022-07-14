package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bartvanbenthem/manifestgen/pkg/convert"
)

var ft, tmpl *string

func ManifestPrinter(input []byte, template string, c convert.Builder) {
	err := c.TemplateParser(input, template)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func main() {

	ft = flag.String("type", "json", "choose input type json / yaml")
	tmpl = flag.String("template", "", "path/file to specific template")
	flag.Parse()

	s := convert.StdinPipeToByte()
	if len(s) == 0 {
		log.Fatal()
	}

	// check if file-type is yaml or json and run corresponding function
	if string(*ft) == string("yaml") {
		ManifestPrinter(s, *tmpl, &convert.YAML{})
	} else if string(*ft) == string("json") {
		ManifestPrinter(s, *tmpl, &convert.JSON{})
	}

}
