package configuration

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Configuration struct {
	NameOfRepository  string `yaml:"NameOfRepository"`
	OwnerOfRepository string `yaml:"OwnerOfRepository"`
}

func LoadConfiguration() (Configuration, error) {

	var config Configuration
	filepath := "configuration/configuration.yaml"

	err := UnmarshalYAMLFile(filepath, &config)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}

func UnmarshalYAMLFile(path string, output interface{}) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, output)
	if err != nil {
		return err
	}

	return nil
}
