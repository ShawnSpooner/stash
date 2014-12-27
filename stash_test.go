package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestDefaultConfigFilePath(t *testing.T) {
	expectedPath := os.Getenv("HOME") + "/.stash"
	path := DefaultConfigPath()
	assert.Equal(t, path, expectedPath)
}

func TestConvertFileToJson(t *testing.T) {
	mockFile := []byte(`{"test":"value"}`)
	json := convertFileToMap(mockFile)
	assert.Equal(t, json["test"], "value")
}

func TestConvertMapToJsonString(t *testing.T) {
	mapFile := map[string]string{"test": "value"}
	jsonValue := convertConfigToJson(mapFile)
	assert.Equal(t, string(jsonValue), `{"test":"value"}`)
}

func TestAddToStash(t *testing.T) {
	stash := Stash{values: map[string]string{}}
	stash.Add("test", "value")
	assert.Equal(t, stash.Get("test"), "value")
}

func TestBuildStashFromReader(t *testing.T) {
	stash := buildStashFromBuffer(strings.NewReader(`{"loaded":"data"}`))
	assert.Equal(t, stash.Get("loaded"), "data")
}
