package python

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// CleanUpVirtualEnv Cleans up a virtual env by deleting .venv/ and __pypackages__/ dirs
func CleanUpVirtualEnv(path string) {
	log.Info(fmt.Sprintf("Teardown: Current working directory: %s", path))

	// cleanup __pypackages__ dir
	if dirExists(filepath.Join(path, "./__pypackages__")) {
		log.Debug("Deleting '__pypackages__/' directory")

		os.RemoveAll(filepath.Join(path, "./__pypackages__"))
	}

	// cleanup virtual env
	if dirExists(filepath.Join(path, "./.venv")) {
		log.Debug("Deleting '.venv/' directory")

		os.RemoveAll(filepath.Join(path, "./.venv"))
	}
}

// SetupVirtualEnv Creates a virtual environment using Poetry, PyFlow, pipenv or pip/venv
func SetupVirtualEnv(path string) {
	log.Info(fmt.Sprintf("Setup: Current working directory: %s", path))

	if fileExists(filepath.Join(path, "pyproject.toml")) {
		// Poetry or PyFlow
		pep518Proc(path)
	} else if fileExists(filepath.Join(path, "Pipfile")) {
		pipenvProc()
	} else if fileExists(filepath.Join(path, "requirements.txt")) {
		pipProc(path)
	}
}
