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
	var stderr bytes.Buffer
	cmd := exec.Command("make", "ui-build")
	cmd.Stderr = &stderr

	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %s", err, stderr.String())
	}

	log.Info("UI assets has been compiled and saved into ./public directory.")

	return nil
}
