package main

import (
	"fmt"

	"github.com/bartvanbenthem/manifestgen/app"
)

func main() {
	var a app.ManifestGenClient
	a.GenerateYamlManifest("testing/values/team.yaml", "testing/templates/team.yaml", "testing/output/team.yaml")
	a.GenerateJSONManifest("testing/values/team.json", "testing/templates/team.json", "testing/output/team.json")
	fmt.Printf("Done\n")
}
