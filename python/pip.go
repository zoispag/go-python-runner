package python

import (
	"fmt"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func pipProc(path string) {
	log.Info("Found 'requirements.txt'. Creating virtual environment using 'pip' & 'venv' module.")

	var cmd *exec.Cmd

	// create virtual env
	cmd = exec.Command("python3", "-m", "venv", ".venv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))

	// install dependencies
	cmd = exec.Command(filepath.Join(path, ".venv/bin/pip"), "install", "-r", "requirements.txt")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))
}
