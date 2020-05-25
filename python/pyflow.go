package python

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func pyflowProc() {
	log.Info("Found 'pyproject.toml'. Creating virtual environment using 'PyFlow'.")

	// install dependencies
	cmd := exec.Command("pyflow", "install")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))
}
