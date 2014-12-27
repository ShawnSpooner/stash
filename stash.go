package main

import (
	"encoding/json"
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

func (s *Stash) SaveStashToWriter(w io.Writer) {
	contents := convertConfigToJson(s.values)
	_, err := w.Write(contents)
	check(err)
}

func (s *Stash) PrettyPrint() {
	for k, v := range s.values {
		println(k, "=>", v)
	}
}

func buildStashFromBuffer(r io.Reader) *Stash {
	config, err := ioutil.ReadAll(r)
	check(err)
	values := convertFileToMap(config)
	return &Stash{values: values}
}

func convertFileToMap(file []byte) map[string]string {
	var config map[string]string
	err := json.Unmarshal(file, &config)
	check(err)
	return config
}

func convertConfigToJson(config map[string]string) []byte {
	jsonConfig, err := json.Marshal(config)
	check(err)
	return jsonConfig
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
