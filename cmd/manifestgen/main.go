package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bartvanbenthem/manifestgen/pkg/convert"
	"github.com/bartvanbenthem/manifestgen/pkg/template"
)

var ft, tmpl, rff, wtf *string

func ManifestPrinter(input []byte, template string, c template.Builder) {
	err := c.TemplateParser(input, template)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func main() {

	ft = flag.String("type", "json", "choose input type json / yaml")
	tmpl = flag.String("template", "", "path/file to specific template")
	rff = flag.String("read-from-file", "", "path/file to specific template")
	wtf = flag.String("write-to-file", "", "path/file to specific template")
	flag.Parse()

	if len(*tmpl) == 0 {
		log.Fatal("Error: Provide path to template file")
	}

	if len(*wtf) == 0 && len(*rff) == 0 {
		s := convert.StdinPipeToByte()
		if len(s) == 0 {
			log.Fatal()
		}

		// check if file-type is yaml or json and run corresponding function
		if string(*ft) == string("yaml") {
			ManifestPrinter(s, *tmpl, &template.YAML{})
		} else if string(*ft) == string("json") {
			ManifestPrinter(s, *tmpl, &template.JSON{})
		}
	} else {
		fmt.Printf("File: %s\n", *wtf)
	}

}
