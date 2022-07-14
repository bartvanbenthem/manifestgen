package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bartvanbenthem/manifestgen/pkg/convert"
	"github.com/bartvanbenthem/manifestgen/pkg/template"
)

var ft, tmpl, rff, wtf *string
var t template.Builder

func ManifestPrinter(data []byte, template string, c template.Builder) {
	err := c.ParseToStdout(data, template)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func ManifestWriter(data []byte, template, file string, c template.Builder) {
	err := c.ParseToFile(data, template, file)
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

	if len(*rff) == 0 {

		s := convert.StdinPipeToByte()
		if len(s) == 0 {
			log.Fatal()
		}

		if len(*wtf) == 0 {
			// check if file-type is yaml or json and run corresponding function
			if string(*ft) == string("yaml") {
				ManifestPrinter(s, *tmpl, &template.YAML{})
			} else if string(*ft) == string("json") {
				ManifestPrinter(s, *tmpl, &template.JSON{})
			}
		} else {
			// check if file-type is yaml or json and run corresponding function
			if string(*ft) == string("yaml") {
				ManifestWriter(s, *tmpl, *wtf, &template.YAML{})
			} else if string(*ft) == string("json") {
				ManifestWriter(s, *tmpl, *wtf, &template.JSON{})
			}
		}
	}

	//!!!!!!!!!!!!!!! create ReadFromFile function in template pkg
	// handle readfrom file errors
	if len(*rff) > 0 {
		if len(*wtf) == 0 {
			// check if file-type is yaml or json and run corresponding function
			if string(*ft) == string("yaml") {
				file, _ := t.ReadFromFile(*ft)
				ManifestPrinter(file, *tmpl, &template.YAML{})
			} else if string(*ft) == string("json") {
				file, _ := t.ReadFromFile(*ft)
				ManifestPrinter(file, *tmpl, &template.JSON{})
			}
		} else {
			// check if file-type is yaml or json and run corresponding function
			if string(*ft) == string("yaml") {
				file, _ := t.ReadFromFile(*ft)
				ManifestWriter(file, *tmpl, *wtf, &template.YAML{})
			} else if string(*ft) == string("json") {
				file, _ := t.ReadFromFile(*ft)
				ManifestWriter(file, *tmpl, *wtf, &template.JSON{})
			}
		}
	}

}
