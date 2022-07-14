package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bartvanbenthem/manifestgen/pkg/convert"
)

func main() {
	jstr := flag.String("string", "", "give json string to marshall into json object")
	file := flag.String("jsonfile", "", "path/file to json file")
	stdinjson := flag.String("json", "", "stdin json input")
	escape := flag.String("escape", "false", "add \\ as escape char")
	srlz := flag.String("serialization", "serialize", "serialize / deserialize")
	flag.Parse()

	if string(*srlz) != "serialize" && string(*srlz) != "deserialize" {
		log.Fatal("Give correct serialization input")
	}

	// deserialize | string-to-json
	if string(*srlz) == "deserialize" {
		output, err := convert.StringToJSON(*jstr)
		if err != nil {
			log.Printf("Error de-serializing: %s", err)
		}

		fmt.Printf("%s\n", output)

	}

	// serialize | json-to-string
	if string(*srlz) == "serialize" {
		var err error
		var content []byte

		if len(string(*file)) != 0 {
			content, err = ioutil.ReadFile(*file)
			if err != nil {
				log.Printf("Error reading file: %s", err)
			}
			output, err := convert.JsonToString(content, convert.StringToBool(*escape))
			if err != nil {
				log.Printf("Error serializing: %s", err)
			}
			fmt.Printf("%s\n", output)
		}

		if len(string(*stdinjson)) != 0 {
			output, err := convert.JsonToString([]byte(*stdinjson), convert.StringToBool(*escape))
			if err != nil {
				log.Printf("Error serializing: %s", err)
			}
			fmt.Printf("%s\n", output)
		}

	}
}
