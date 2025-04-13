package utils

import (
	"fmt"
	"os/exec"
)

func TODO(str string) error {
	return fmt.Errorf("TODO: %s", str)
}

func IsToolExists(toolName string) bool {
	_, err := exec.LookPath(toolName)
	if err != nil {
		return false
	}

	return true
}
