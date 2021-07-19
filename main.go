package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Printf("manifestgen\n")
}

type ManifestGenClient struct{}

func (c *ManifestGenClient) ConvertJSON2Yaml(pathJSONFile string) {
}

func (c *ManifestGenClient) ConvertYaml2JSON(pathYamlFile string) {
}

func (c *ManifestGenClient) GenerateJSONManifest(pathValuesFile, pathTemplateFile, pathOutputFolder string) {
}

// Function to generate the configurable manifest from the values and template files
// pathValuesFile wants a string containing path to the values yaml file
// pathTemplateFile wants a string containing path to the template yaml file
// pathOutputFolder wants a string containing path to the output folder
func (c *ManifestGenClient) GenerateYamlManifest(pathValuesFile, pathTemplateFile, pathOutputFolder string) {
	// open the values yaml file
	yamlFile, err := os.Open(pathValuesFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	defer yamlFile.Close()

	// read the values file and create a byte slice output
	byteValue, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	// create a map to store the unmarshalled byte slice
	var values map[string]interface{}

	// unmarshal byte slice into the values map
	err = yaml.Unmarshal(byteValue, &values)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	tpl, err := template.ParseFiles(pathTemplateFile)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	valueFile := strings.Split(pathValuesFile, "/")
	fileName := strings.Split(fmt.Sprintf(valueFile[len(valueFile)-1]), ".")
	genFile, err := os.Create(fmt.Sprintf("%v/%v.yaml", pathOutputFolder, fileName[0]))
	if err != nil {
		log.Fatalln(err)
	}
	defer genFile.Close()

	// Execute template var injection and write to file
	err = tpl.Execute(genFile, values)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

}

type Files struct {
	FileName    string
	FilePath    string
	PathAndFile string
}

func ReadFiles(path string) ([]Files, error) {
	var err error
	var file Files
	var files []Files

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range fs {
		file.FileName = f.Name()
		file.FilePath = path
		file.PathAndFile = fmt.Sprintf("%v/%v", path, f.Name())
		files = append(files, file)
	}

	return files, err
}

func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
