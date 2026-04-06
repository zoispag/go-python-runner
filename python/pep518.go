package python

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func pep518Proc(path string) {
	tomlContent := readPyProjectToml(path)

	if isUv(path) {
		uvProc(path)
	} else if isPoetry(tomlContent) {
		poetryProc(path)
	}
}

func readPyProjectToml(path string) string {
	// read the whole file at once
	b, err := ioutil.ReadFile(filepath.Join(path, "pyproject.toml"))
	if err != nil {
		panic(err)
	}

	return string(b)
}

func isPoetry(tomlContent string) bool {
	return strings.Contains(tomlContent, "tool.poetry")
}

func isUv(path string) bool {
	return fileExists(filepath.Join(path, "uv.lock"))
}
