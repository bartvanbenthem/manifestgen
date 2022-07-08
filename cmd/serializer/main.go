package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func StringPointerToBool(s *string) bool {
	i, err := strconv.ParseBool(*s)
	if err != nil {
		log.Fatal(err)
	}

	return bool(i)
}

func jsonToString(file []byte, escape bool) (string, error) {
	var output []byte
	var err error

	//fmt.Println(string(content))
	var data map[string]any
	err = json.Unmarshal(file, &data)
	if err != nil {
		return string(output), err
	}

	output, err = json.Marshal(&data)

	if escape == true {
		//fmt.Println(strings.Replace(string(output), "\"", "\\\"", -1))
		return strings.Replace(string(output), "\"", "\\\"", -1), nil

	} else {
		//fmt.Println(string(output))
		return string(output), nil
	}
}

func stringToJSON(jsonstr string) (string, error) {
	var output []byte
	var err error
	var body map[string]interface{}

	err = json.Unmarshal([]byte(*&jsonstr), &body)
	if err != nil {
		//fmt.Printf("Error Unmarshalling: %s", err)
		return string(output), err
	}

	output, err = json.Marshal(body)
	if err != nil {
		//fmt.Printf("Error Marshalling: %s", err)
		return string(output), err
	}

	return string(output), err
}

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

	// serialize | string-to-json
	if string(*srlz) == "serialize" {
		output, err := stringToJSON(*jstr)
		if err != nil {
			log.Printf("Error serializing: %s", err)
		}

		fmt.Printf("%s\n", output)

	}

	// de-serialize | json-to-string
	if string(*srlz) == "deserialize" {
		var err error
		var content []byte

		if len(string(*file)) != 0 {
			content, err = ioutil.ReadFile(*file)
			if err != nil {
				log.Printf("Error reading file: %s", err)
			}
			output, err := jsonToString(content, StringPointerToBool(escape))
			if err != nil {
				log.Printf("Error de-serializing: %s", err)
			}
			fmt.Printf("%s\n", output)
		}

		if len(string(*stdinjson)) != 0 {
			output, err := jsonToString([]byte(*stdinjson), StringPointerToBool(escape))
			if err != nil {
				log.Printf("Error de-serializing: %s", err)
			}
			fmt.Printf("%s\n", output)
		}

	}
}
