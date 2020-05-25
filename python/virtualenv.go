package python

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// CleanUpVirtualEnv Cleans up a virtual env by deleting .venv/ and __pypackages__/ dirs
func CleanUpVirtualEnv() {
	// cleanup __pypackages__ dir
	if dirExists("./__pypackages__") {
		log.Debug("Deleting '__pypackages__/' directory")

		os.RemoveAll("./__pypackages__")
	}

	// cleanup virtual env
	if dirExists("./.venv") {
		log.Debug("Deleting '.venv/' directory")

		os.RemoveAll("./.venv")
	}
}

// SetupVirtualEnv Creates a virtual environment using Poetry, PyFlow, pipenv or pip/venv
func SetupVirtualEnv() {
	if fileExists("pyproject.toml") {
		if isPoetry() {
			poetryProc()
		} else if isPyFlow() {
			pyflowProc()
		}
	} else if fileExists("Pipfile") {
		pipenvProc()
	} else if fileExists("requirements.txt") {
		pipProc()
	}
}
