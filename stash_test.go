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
	json, _ := convertFileToMap(mockFile)
	assert.Equal(t, json["test"], "value")
}

func TestErrorReturnedWithInvalidJson(t *testing.T) {
	mockFile := []byte(`{"test" ->broken"}`)
	_, err := convertFileToMap(mockFile)
	assert.NotNil(t, err)
}

func TestConvertMapToJsonString(t *testing.T) {
	mapFile := map[string]string{"test": "value"}
	jsonValue, _ := convertConfigToJson(mapFile)
	assert.Equal(t, string(jsonValue), `{"test":"value"}`)
}

func TestAddToStash(t *testing.T) {
	stash := Stash{values: map[string]string{}}
	stash.Add("test", "value")
	assert.Equal(t, stash.Get("test"), "value")
}

func TestBuildStashFromReader(t *testing.T) {
	stash, _ := buildStashFromBuffer(strings.NewReader(`{"loaded":"data"}`))
	assert.Equal(t, stash.Get("loaded"), "data")
}

func TestFormatOutput(t *testing.T) {
	stash := Stash{values: map[string]string{"test": "value"}}
	output := stash.Format()
	assert.Equal(t, output, "test => value\n")
}
