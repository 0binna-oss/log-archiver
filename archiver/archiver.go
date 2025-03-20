package archiver

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ArchiveLogs archives the given log file to the specified directory with a time stamp
func ArchiveLogs(logFilePath, archiveDir string) (string, error) {
	// Open the log file
	file, err := os.Open(logFilePath)
	if err != nil {
		return "", fmt.Errorf("error opening log file: %w", err)
	}
	defer file.Close()

	// Create the archive directory if it doesn't exist
	if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
		if err := os.Mkdir(archiveDir, 0755); err != nil {
			return "", fmt.Errorf("error creating archive directory: %w", err)
		}
	}

	// Generate a timestamp for the archive file
	timestamp := time.Now().Format("2025-03-20_07-46-05")
	archiveFilename := fmt.Sprintf("log_%s.txt", timestamp)
	archiveFilePath := filepath.Join(archiveDir, archiveFilename)

	// Create the archive file
	archiveFile, err := os.Create(archiveFilePath)
	if err != nil {
		return "", fmt.Errorf("error creating archive file: %w", err)
	}
	defer archiveFile.Close()

	// Read the log file line by line and write to the archive file
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(archiveFile)
	for scanner.Scan() {
		line := scanner.Text()
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return "", fmt.Errorf("error writing to archive file: %w", err)
		}
	}
	writer.Flush()

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading log file: %w", err)
	}

	return archiveFilePath, nil
}
