package python

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// CleanUpVirtualEnv Cleans up a virtual env by deleting .venv/ dir
func CleanUpVirtualEnv() {
	// cleanup virtual env
	os.RemoveAll("./.venv")
}

// SetupVirtualEnv Creates a virtual environment using Poetry, pipenv or pip/venv
func SetupVirtualEnv() {
	if fileExists("pyproject.toml") {
		poetryProc()
	} else if fileExists("Pipfile") {
		pipenvProc()
	} else if fileExists("requirements.txt") {
		pipProc()
	}
}

func poetryProc() {
	log.Info("Found 'pyproject.toml'. Creating virtual environment using 'Poetry'.")

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
	log.Info("Found 'Pipfile'. Creating virtual environment using 'pipenv'.")

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
	log.Info("Found 'requirements.txt'. Creating virtual environment using 'pip' & 'venv' module.")

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
