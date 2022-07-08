package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	str := flag.String("jsonfile", "", "path/file to json file")
	escape := flag.String("escape", "false", "add \\ as escape char")
	flag.Parse()

	content, err := ioutil.ReadFile(*str)
	if err != nil {
		fmt.Println("Err")
	}

	//fmt.Println(string(content))

	var data map[string]any
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error Marshalling: %s", err)
	}

	output, err := json.Marshal(&data)

	if string(*escape) == "true" {
		fmt.Println(strings.Replace(string(output), "\"", "\\\"", -1))
	} else {
		fmt.Println(string(output))
	}

}
