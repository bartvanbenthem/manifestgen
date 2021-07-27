package app

type Converter interface {
	Convert(pathInputFile, pathOutputFile string)
}

type YAML2JSON struct {
}

func (c YAML2JSON) Convert(pathYamlFile, pathJsonFile string) {}

type JSON2YAML struct {
}

func (c JSON2YAML) Convert(pathYamlFile, pathJsonFile string) {}
