package ui

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CompileJS() (string, error) {
	var stderr bytes.Buffer
	cmd := exec.Command("make", "ui-build-serverside")
	cmd.Stderr = &stderr

	source, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, stderr.String())
	}

	return string(source), nil
}

func BuildUI() error {
	log.Info("Detected changes, rebuilding the UI.")

	var stderr bytes.Buffer
	cmd := exec.Command("make", "ui-build")
	cmd.Stderr = &stderr

	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %s", err, stderr.String())
	}

	return nil
}
