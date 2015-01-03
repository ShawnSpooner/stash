package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Stash struct {
	values map[string]string
	path   string
}

func DefaultConfigPath() string {
	return os.Getenv("HOME") + "/.stash"
}

func (s *Stash) Add(key string, value string) {
	s.values[key] = value
}

func (s *Stash) Get(key string) string {
	return s.values[key]
}

//Writes the stashes contents encoded as JSON to the writer supplied
func (s *Stash) SaveStashToWriter(w io.Writer) error {
	contents, err := convertConfigToJson(s.values)
	_, err = w.Write(contents)
	return err
}

//Returns a pretty formatted string of the contents of the stash
//Keys => Values
func (s *Stash) Format() string {
	var output string
	for k, v := range s.values {
		output += fmt.Sprintf("%v => %v\n", k, v)
	}
	return output
}

//Builds a new stash from the Reader supplied to it by reading from the reader
//and converting the JSON into a map[string]string. It then initilizes a new stash around
//the map
func buildStashFromBuffer(r io.Reader) (*Stash, error) {
	config, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	values, err := convertFileToMap(config)
	return &Stash{values: values}, err
}

//Convert the supplied byte array(containing valid JSON) into a map
func convertFileToMap(file []byte) (map[string]string, error) {
	var config map[string]string
	err := json.Unmarshal(file, &config)
	return config, err
}

//Convert the supplied map to a byte array containing the map marshalled as JSON
func convertConfigToJson(config map[string]string) ([]byte, error) {
	return json.Marshal(config)
}
