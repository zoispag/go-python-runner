package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	python "github.com/zoispag/go-python-runner/python"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Set debug mode on
	if os.Getenv("GO_DEBUG_MODE") == "on" {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Error(fmt.Sprintf("Job resulted in error: %s", err.Error()))
	}
	python.CleanUpVirtualEnv(path)
	python.SetupVirtualEnv(path)

	cmd := python.GetPythonRunCommand(path, "script.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("Job resulted in error: %s", err.Error()))
	}
	log.Info(string(out))
}
