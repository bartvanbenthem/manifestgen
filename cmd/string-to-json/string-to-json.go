package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

func main() {
	input := flag.String("string", "", "give json string to marshall into json object")
	flag.Parse()

	var body map[string]interface{}

	err := json.Unmarshal([]byte(*input), &body)
	if err != nil {
		fmt.Printf("Error Unmarshalling: %s", err)
	}

	output, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Error Marshalling: %s", err)
	}

	fmt.Println(string(output))
}
