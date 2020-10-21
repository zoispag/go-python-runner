package python

import (
	"os/exec"
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

// GetPythonRunCommand Returns an *exec.Cmd to be handled with the provided "scriptName.py" using the python binary from virtual environment
func GetPythonRunCommand(scriptName string) *exec.Cmd {
	pythonPath := getPythonCommand()
	return exec.Command(pythonPath, scriptName)
}

// ExecutePython Executes the provided "scriptName.py" using the python binary from virtual environment
func ExecutePython(scriptName string) ([]byte, error) {
	// run python job
	pythonPath := getPythonCommand()
	cmd := exec.Command(pythonPath, scriptName)
	return cmd.CombinedOutput()
}
