package convert

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func StringToBool(s string) bool {
	i, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}

	return bool(i)
}

func StringToInt32(s string) int32 {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return int32(i)
}

func StringToJSON(jsonstr string) (string, error) {
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

func JsonToString(file []byte, escape bool) (string, error) {
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
