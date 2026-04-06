package python

import (
	"os/exec"
	"path/filepath"
)

func getPythonCommand(path string) string {
	if dirExists(filepath.Join(path, "./.venv")) {
		return filepath.Join(path, ".venv/bin/python")
	}
	return "python"
}

// GetPythonRunCommand Returns an *exec.Cmd to be handled with the provided "scriptName.py" using the python binary from virtual environment
func GetPythonRunCommand(path string, scriptName string) *exec.Cmd {
	return exec.Command(getPythonCommand(path), scriptName)
}

// ExecutePython Executes the provided "scriptName.py" using the python binary from virtual environment
func ExecutePython(path string, scriptName string) ([]byte, error) {
	return GetPythonRunCommand(path, scriptName).CombinedOutput()
}
