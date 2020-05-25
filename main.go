package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	python "github.com/zoispag/go-python-runner/python"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func main() {
	python.CleanUpVirtualEnv()
	python.SetupVirtualEnv()

	python.ExecutePython("script.py")
}
