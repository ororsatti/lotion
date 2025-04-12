package state

import (
	"errors"
	"os"
)

func isDirExists(path string) bool {
	stat, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) || !stat.IsDir() {
		return false
	}

	return true
}

func isFileExists(path string) bool {
	stat, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) || stat.IsDir() {
		return false
	}

	return true
}
