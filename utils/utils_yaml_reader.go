package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigEntry struct {
	Name string `yaml:"name"`
	FilePath string `yaml:"file_path"`
	Format string `yaml:"format"`
	BacklogSize int64 `yaml:"backlog_size"`
}

func RetrieveParsedYaml(yamlFilePath string) []ConfigEntry {
	fmt.Println(yamlFilePath)
	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		panic(err)
	}
	var config []ConfigEntry
	fmt.Println(yamlFilePath)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	return config
}