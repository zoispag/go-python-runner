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

func poetryProc() {
	// Make sure .venv will be in project
	if !fileExists("Pipfile") {
		cmd := exec.Command("poetry", "config", "--local", "virtualenvs.in-project", "true")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(fmt.Sprintf("%s", err.Error()))
		}
		log.Debug(string(out))
	}

	// install dependencies
	cmd := exec.Command("poetry", "install")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))
}

func pipenvProc() {
	// Make sure .venv will be in project
	os.Setenv("PIPENV_VENV_IN_PROJECT", "1")

	// install dependencies
	cmd := exec.Command("pipenv", "install")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))
}

func pipProc() {
	var cmd *exec.Cmd

	// create virtual env
	cmd = exec.Command("python3", "-m", "venv", ".venv")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))

	// install dependencies
	cmd = exec.Command("./.venv/bin/pip", "install", "-r", "requirements.txt")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Debug(string(out))
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func main() {
	// cleanup virtual env
	os.RemoveAll("./.venv")

	var cmd *exec.Cmd
	if fileExists("pyproject.toml") {
		log.Info("Found 'pyproject.toml'. Creating virtual environment using 'Poetry'.")
		poetryProc()
	} else if fileExists("Pipfile") {
		log.Info("Found 'Pipfile'. Creating virtual environment using 'pipenv'.")
		pipenvProc()
	} else if fileExists("requirements.txt") {
		log.Info("Found 'requirements.txt'. Creating virtual environment using 'pip' & 'venv' module.")
		pipProc()
	}

	// run python job
	cmd = exec.Command("./.venv/bin/python", "script.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(fmt.Sprintf("%s", err.Error()))
	}
	log.Info(string(out))
}
