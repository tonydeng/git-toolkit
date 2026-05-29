package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Meta struct {
	Verify map[string][]string
	Gitlab map[string]string
}

func NewReadConfig(filename string) map[string][]string {
	var meta Meta
	data, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(data, &meta)
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
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
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	err = yaml.Unmarshal(data, &meta)
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	return meta.Gitlab
}
