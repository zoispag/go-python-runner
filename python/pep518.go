package python

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func pep518Proc(path string) {
	tomlContent := readPyProjectToml(path)
	if tomlContent == "" {
		return
	}

	if isUv(path) {
		uvProc(path)
	} else if isPoetry(tomlContent) {
		poetryProc(path)
	}
}

func readPyProjectToml(path string) string {
	b, err := os.ReadFile(filepath.Join(path, "pyproject.toml"))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return string(b)
}

func isPoetry(tomlContent string) bool {
	return strings.Contains(tomlContent, "tool.poetry")
}

func isUv(path string) bool {
	return fileExists(filepath.Join(path, "uv.lock"))
}
