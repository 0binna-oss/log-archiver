package rotator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// RotateLogs deletes old logsbased on sized and age
func RotateLogs(archiveDir string, maxLogSizeMB, maxLogAgeDays int) error {
	files, err := os.ReadDir(archiveDir)
	if err != nil {
		return fmt.Errorf("error reading archive directory: %w", err)
	}

	for _, file := range files {
		filePath := filepath.Join(archiveDir, file.Name())
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return fmt.Errorf("error getting file info: %w", err)
		}

		// Check file sizes
		if maxLogAgeDays > 0 && time.Since(fileInfo.ModTime()) > time.Duration(maxLogAgeDays)*24*time.Hour {
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("error deleting file: %w", err)
			}
		}
	}

	return nil
}
