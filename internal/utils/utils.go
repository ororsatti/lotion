package utils

import (
	"fmt"
)

func TODO(str string) error {
	return fmt.Errorf("TODO: %s", str)
}
