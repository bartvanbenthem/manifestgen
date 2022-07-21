package convert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToBool(t *testing.T) {
	b := false
	v := StringToBool("false")

	if v == b {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not string, but is %T", v))
	}

}

func TestStringToInt32(t *testing.T) {
	b := int32(33)
	v := StringToInt32("33")

	if v == b {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not int32, but is %T", v))
	}

}

func TestStringToJSON(t *testing.T) {
	v, err := StringToJSON("{\"json\": \"test\"}")
	if err != nil {
		t.Error("failed StringToJSON()")
	}

	if reflect.TypeOf(v) == reflect.TypeOf([]byte("test")) {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not []byte, but is %T", v))
	}

}

func TestJsonToString(t *testing.T) {
	b := true
	v, err := JsonToString([]byte("{\"json\": \"test\"}"), b)
	if err != nil {
		t.Error("failed JsonToString()")
	}

	if reflect.TypeOf(v) == reflect.TypeOf(string("test")) {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not string, but is %T", v))
	}

}

func TestJsonToYaml(t *testing.T) {
	v, err := JsonToYaml([]byte("{\"json\": \"test\"}"))
	if err != nil {
		t.Error("failed JsonToYaml()")
	}

	if reflect.TypeOf(v) == reflect.TypeOf([]byte("{yaml: test}")) {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not []byte, but is %T", v))
	}

}

func TestStdinPipeToByte(t *testing.T) {

	v := StdinPipeToByte()
	if reflect.TypeOf(v) == reflect.TypeOf([]byte("test")) {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not []byte, but is %T", v))
	}

}
