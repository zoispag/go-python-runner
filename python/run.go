package python

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func getPythonCommand() string {
	if dirExists("./__pypackages__") {
		return "pyflow"
	} else if dirExists("./.venv") {
		return "./.venv/bin/python"
	} else {
		return "python"
	}
}

// ExecutePython Executes the provided "scriptName.py" using the python binary from virtual environment
func ExecutePython(scriptName string) {
	// run python job
	pythonPath := getPythonCommand()
	cmd := exec.Command(pythonPath, scriptName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Info(string(out))
}
