package python

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func uvProc(path string) {
	log.Info("Found 'uv.lock'. Creating virtual environment using 'uv'.")

	cmd := exec.Command("uv", "sync")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug(string(out))
}
