package python

import (
	"io/ioutil"
	"strings"
)

func pep518Proc() {
	tomlContent := readPyProjectToml()

	if isPoetry(tomlContent) {
		poetryProc()
	} else if isPyFlow(tomlContent) {
		pyflowProc()
	}
}

func readPyProjectToml() string {
	// read the whole file at once
	b, err := ioutil.ReadFile("pyproject.toml")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func isPoetry(tomlContent string) bool {
	return strings.Contains(tomlContent, "tool.poetry")
}

func isPyFlow(tomlContent string) bool {
	return strings.Contains(tomlContent, "tool.pyflow")
}
