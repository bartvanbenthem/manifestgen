package template

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	ioutil.WriteFile("../../project/testdata/template",
		[]byte("{{.test}}"), 0644)
	v, err := ReadFromFile("../../project/testdata/template")
	if err != nil {
		t.Error("failed ReadFromFile()")
	}

	if reflect.TypeOf(v) == reflect.TypeOf([]byte("test")) {
		// do nothing
	} else {
		t.Error(fmt.Sprintf("type is not []byte, but is %T", v))
	}

}

// !! fin out how to proper test std out

func TestParseToStdout(t *testing.T) {
	// test the functions with a JSON receiver
	j := JSON{}
	err := j.ParseToStdout([]byte("{\"test\": \"json\"}"),
		"../../project/testdata/template")
	if err != nil {
		t.Error("failed JSON ParseToStdout()")
	}
	// test the functions with a YAML receiver
	y := YAML{}
	err = y.ParseToStdout([]byte("{test: yaml}"),
		"../../project/testdata/template")
	if err != nil {
		t.Error("failed YAML ParseToStdout()")
	}

}

func TestParseToFile(t *testing.T) {
	j := JSON{}
	err := j.ParseToFile([]byte("{\"test\": \"json\"}"),
		"../../project/testdata/template",
		"../../project/testdata/output_test_json")
	if err != nil {
		t.Error("failed JSON ParseToStdout()")
	}

	y := YAML{}
	err = y.ParseToFile([]byte("{test: yaml}"),
		"../../project/testdata/template",
		"../../project/testdata/output_test_yaml")
	if err != nil {
		t.Error("failed YAML ParseToStdout()")
	}
}

func testBuilderInterface(t *testing.T) {
	func(b Builder) {
	}(&TestBuilder{})

}
