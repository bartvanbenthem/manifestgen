package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bartvanbenthem/manifestgen/internal/template"
	"github.com/bartvanbenthem/manifestgen/pkg/convert"
)

var ft, tmpl, rff, wtf *string

func ManifestPrinter(data []byte, template string, c template.Builder) {
	err := c.ParseToStdout(data, template)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func ManifestWriter(data []byte, template, file string, c template.Builder) {
	err := c.ParseToFile(data, template, file)
	if err != nil {
		log.Printf("Error during %v manifest generation: %s\n", file, err)
	}

	log.Printf("%v manifest is generated", file)
}

func main() {

	ft = flag.String("type", "json", "choose input type json / yaml")
	tmpl = flag.String("template", "", "path/file to template file")
	rff = flag.String("read-from-file", "", "path/file to variables file")
	wtf = flag.String("write-to-file", "", "path/file for writing generated file")
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
	// handle read from file errors
	if len(*rff) > 0 {
		if len(*wtf) == 0 {
			// check if file-type is yaml or json and run corresponding function
			if string(*ft) == string("yaml") {
				file, err := template.ReadFromFile(*rff)
				if err != nil {
					log.Fatal(err)
				}
				ManifestPrinter(file, *tmpl, &template.YAML{})
			} else if string(*ft) == string("json") {
				file, err := template.ReadFromFile(*rff)
				if err != nil {
					log.Fatal(err)
				}
				ManifestPrinter(file, *tmpl, &template.JSON{})
			}
		} else {
			// check if file-type is yaml or json and run corresponding function
			if string(*ft) == string("yaml") {
				file, err := template.ReadFromFile(*rff)
				if err != nil {
					log.Fatal(err)
				}
				ManifestWriter(file, *tmpl, *wtf, &template.YAML{})
			} else if string(*ft) == string("json") {
				file, err := template.ReadFromFile(*rff)
				if err != nil {
					log.Fatal(err)
				}
				ManifestWriter(file, *tmpl, *wtf, &template.JSON{})
			}
		}
	}

}
