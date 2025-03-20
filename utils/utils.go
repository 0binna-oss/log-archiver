package utils

import (
	"fmt"
	"os"
)

// ValidateFilePath checks if a file exists and is accessible
func ValidateFilePath(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}
	return nil
}
