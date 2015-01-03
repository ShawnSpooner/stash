package main

import (
	"github.com/atotto/clipboard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCopiesTheValueToTheClipboard(t *testing.T) {
	stash := &Stash{values: map[string]string{"test": "value"}}
	GetEntry(stash, []string{"test"})
	clipping, _ := clipboard.ReadAll()
	assert.Equal(t, clipping, "value")
}
