package main

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func main() {
	// cleanup virtual env
	os.RemoveAll("./.venv")

	setupVirtualEnv()

	// run python job
	cmd := exec.Command("./.venv/bin/python", "script.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Info(string(out))
}
