package convert

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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

func JsonToYaml(serialized []byte) ([]byte, error) {
	var output []byte
	var err error
	var body map[string]interface{}

	err = json.Unmarshal(serialized, &body)
	if err != nil {
		return output, err
	}

	output, err = yaml.Marshal(body)
	if err != nil {
		return output, err
	}

	return output, err

}

func StdinPipeToByte() []byte {
	var out []string
	// create a reader from the standard input.
	r := bufio.NewReader(os.Stdin)
	// create a buffer of 4KB.
	buf := make([]byte, 0, 4*1024)

	for {
		// read data from the standard input into the buffer.
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			// check for errors
			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		// add contents of the buffer to the slice of string.
		out = append(out, fmt.Sprintf("%s", string(buf)))
		// check for errors
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	// return []byte by converting []string to a single string
	return []byte(strings.Join(out, " "))
}
