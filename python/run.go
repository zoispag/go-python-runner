package python

import (
	"os/exec"
	"path/filepath"
)

func getPythonCommand(path string) string {
	if dirExists(filepath.Join(path, "./__pypackages__")) {
		return "pyflow"
	} else if dirExists(filepath.Join(path, "./.venv")) {
		return filepath.Join(path, ".venv/bin/python")
	} else {
		return "python"
	}
}

// GetPythonRunCommand Returns an *exec.Cmd to be handled with the provided "scriptName.py" using the python binary from virtual environment
func GetPythonRunCommand(path string, scriptName string) *exec.Cmd {
	pythonPath := getPythonCommand(path)
	return exec.Command(pythonPath, scriptName)
}

// ExecutePython Executes the provided "scriptName.py" using the python binary from virtual environment
func ExecutePython(path string, scriptName string) ([]byte, error) {
	// run python job
	pythonPath := getPythonCommand(path)
	cmd := exec.Command(pythonPath, scriptName)
	return cmd.CombinedOutput()
}
