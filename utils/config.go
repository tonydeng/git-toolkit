package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Meta struct {
	Verify map[string][]string
	Gitlab map[string]string
}

func NewReadConfig(filename string) map[string][]string {
	var meta Meta
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Fatal error reading config file: %s \n", err)
		return nil
	}
	if err := yaml.Unmarshal(data, &meta); err != nil {
		fmt.Printf("Fatal error parsing config file: %s \n", err)
		return nil
	}
	return meta.Verify
}

func GetAllKey(m map[string][]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetKey(m map[string][]string, key string) bool {
	for elem := range m {
		if elem == key {
			return true
		}
	}
	return false
}

func ReadGitlabConfig(filename string) map[string]string {
	var meta Meta
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Fatal error reading config file: %s \n", err)
		return nil
	}
	if err := yaml.Unmarshal(data, &meta); err != nil {
		fmt.Printf("Fatal error parsing config file: %s \n", err)
		return nil
	}
	return meta.Gitlab
}
