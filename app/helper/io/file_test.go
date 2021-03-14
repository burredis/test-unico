package io

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var filename = "xpto.txt"

func TestWritefile(t *testing.T) {
	Writefile(filename, "xpto")
	content, _ := ioutil.ReadFile(filename)
	assert.GreaterOrEqual(t, len(string(content)), 1)
}

func TestRemovefile(t *testing.T) {
	Removefile(filename)
	content, _ := ioutil.ReadFile(filename)
	assert.LessOrEqual(t, len(string(content)), 0)
}
