package main

import (
	"testing"

	"github.com/bartvanbenthem/manifestgen/internal/template"
)

func TestManifestPrinter(t *testing.T) {
	ManifestPrinter([]byte("{\"test\": \"json\"}"),
		"../../project/testdata/template",
		&template.TestBuilder{})
}

func TestManifestWriter(t *testing.T) {
	ManifestWriter([]byte("{\"test\": \"json\"}"),
		"../../project/testdata/template",
		"../../project/testdata/testfile",
		&template.TestBuilder{})
}
