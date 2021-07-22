package app

type ConverterClient struct {
}

type Converter interface {
	Convert(pathInputFile, pathOutputFile string)
}

func (*ConverterClient) ConvertYamltoJSON(pathYamlFile, pathJsonFile string) {}
func (*ConverterClient) ConvertJSONtoYaml(pathYamlFile, pathJsonFile string) {}
