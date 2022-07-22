package convert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bartvanbenthem/manifestgen/internal/assert"
)

func TestStringToBool(t *testing.T) {
	// table test struct
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{
			name: "False",
			val:  "false",
			want: false,
		},
		{
			name: "True",
			val:  "true",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := StringToBool(tt.val)
			// compare the expected and actual values.
			assert.Equal(t, v, tt.want)
		})
	}
}

func TestStringToInt32(t *testing.T) {
	b := int32(33)
	v := StringToInt32("33")

	assert.Equal(t, v, b)
}

func TestStringToJSON(t *testing.T) {
	v, err := StringToJSON("{\"json\": \"test\"}")
	if err != nil {
		t.Error("failed StringToJSON()")
	}

	if reflect.TypeOf(v) != reflect.TypeOf([]byte("test")) {
		t.Error(fmt.Sprintf("type is not []byte, but is %T", v))
	}

}

func TestJsonToString(t *testing.T) {
	b := true
	v, err := JsonToString([]byte("{\"json\": \"test\"}"), b)
	if err != nil {
		t.Error("failed JsonToString()")
	}

	assert.StringContains(t, v, "test")

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
