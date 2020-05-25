package python

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// ExecutePython Executes the provided "scriptName.py" using the python binary from virtual environment
func ExecutePython(scriptName string) {
	// run python job
	cmd := exec.Command("./.venv/bin/python", scriptName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Info(string(out))
}
