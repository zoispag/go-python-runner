package python

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func pipenvProc() {
	log.Info("Found 'Pipfile'. Creating virtual environment using 'pipenv'.")

	// Make sure .venv will be in project
	os.Setenv("PIPENV_VENV_IN_PROJECT", "1")

	// install dependencies
	cmd := exec.Command("pipenv", "install")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))
}
