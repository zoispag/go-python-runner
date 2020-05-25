package python

import (
	"io/ioutil"
	"strings"
)

func readPyProjectToml() string {
	// read the whole file at once
	b, err := ioutil.ReadFile("pyproject.toml")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func isPoetry() bool {
	return strings.Contains(readPyProjectToml(), "tool.poetry")
}

func isPyFlow() bool {
	return strings.Contains(readPyProjectToml(), "tool.pyflow")
}
