package utils

import (
	"fmt"
	"testing"
)

func TestNewReadConfig(t *testing.T) {
	m := NewReadConfig("../config.yaml")
	fmt.Println(m["github.com"][0])
	if m["github.com"][0] != "example.com" {
		t.Errorf("TestNewReadConfig: something is wrong with {%s}", m["github.com"][0])
	}
}

func TestGetAllKey(t *testing.T) {
	m := make(map[string][]string)
	m["digital"] = []string{"1", "2"}
	m["letter"] = []string{"a", "b"}
	keys := GetAllKey(m)
	fmt.Println(keys)
	if !IsContain(keys, "digital") || !IsContain(keys, "letter") {
		t.Errorf("TestGetAllKey: something is wrong with {%s}", keys)
	}
}

func TestGetKey(t *testing.T) {
	m := make(map[string][]string)
	m["digital"] = []string{"1", "2"}
	m["letter"] = []string{"a", "b"}
	fmt.Println(GetKey(m, "digital"))
	if !GetKey(m, "digital") {
		t.Errorf("TestGetKey: something is wrong with {%t}", GetKey(m, "digital"))
	}
}
