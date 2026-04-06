package python

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// CleanUpVirtualEnv Cleans up a virtual env by deleting .venv/ dir
func CleanUpVirtualEnv(path string) {
	log.Info(fmt.Sprintf("Teardown: Current working directory: %s", path))

	// cleanup virtual env
	if dirExists(filepath.Join(path, "./.venv")) {
		log.Debug("Deleting '.venv/' directory")

		os.RemoveAll(filepath.Join(path, "./.venv"))
	}
}

// SetupVirtualEnv Creates a virtual environment using uv, Poetry, pipenv or pip/venv
func SetupVirtualEnv(path string) {
	log.Info(fmt.Sprintf("Setup: Current working directory: %s", path))

	if fileExists(filepath.Join(path, "pyproject.toml")) {
		// uv or Poetry
		pep518Proc(path)
	} else if fileExists(filepath.Join(path, "Pipfile")) {
		pipenvProc(path)
	} else if fileExists(filepath.Join(path, "requirements.txt")) {
		pipProc(path)
	}
}
