package python

import (
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func pipProc(path string) {
	log.Info("Found 'requirements.txt'. Creating virtual environment using 'pip' & 'venv' module.")

	cmd := exec.Command("python3", "-m", "venv", ".venv")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug(string(out))

	cmd = exec.Command(filepath.Join(path, ".venv/bin/pip"), "install", "-r", "requirements.txt")
	cmd.Dir = path
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug(string(out))
}
