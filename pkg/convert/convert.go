package convert

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
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

func StringToJSON(jsonstr string) ([]byte, error) {
	var output []byte
	var err error
	var body map[string]interface{}

	err = json.Unmarshal([]byte(*&jsonstr), &body)
	if err != nil {
		return output, err
	}

	output, err = json.Marshal(body)
	if err != nil {
		return output, err
	}

	return output, err
}

func JsonToString(file []byte, escape bool) (string, error) {
	var output []byte
	var err error
	var data map[string]any

	err = json.Unmarshal(file, &data)
	if err != nil {
		return string(output), err
	}

	output, err = json.Marshal(&data)

	if escape == true {
		return strings.Replace(string(output), "\"", "\\\"", -1), nil

	} else {
		return string(output), nil
	}
}

func JsonToYaml(serialized string) ([]byte, error) {
	var output []byte
	var err error
	var body map[string]interface{}

	err = json.Unmarshal([]byte(*&serialized), &body)
	if err != nil {
		return output, err
	}

	output, err = yaml.Marshal(body)
	if err != nil {
		return output, err
	}

	return output, err

}
