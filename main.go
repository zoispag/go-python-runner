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
	python.CleanUpVirtualEnv()
	python.SetupVirtualEnv()

	cmd := python.GetPythonRunCommand("script.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("Job resulted in error: %s", err.Error()))
	}
	log.Info(string(out))
}
