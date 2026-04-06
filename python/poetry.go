package python

import (
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func poetryProc(path string) {
	log.Info("Found 'pyproject.toml'. Creating virtual environment using 'Poetry'.")

	// Make sure .venv will be in project
	if !fileExists(filepath.Join(path, "poetry.toml")) {
		cmd := exec.Command("poetry", "config", "--local", "virtualenvs.in-project", "true")
		cmd.Dir = path
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(err.Error())
		}
		log.Debug(string(out))
	}

	// install dependencies
	cmd := exec.Command("poetry", "install")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug(string(out))
}
