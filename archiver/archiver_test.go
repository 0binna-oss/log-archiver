package archiver

import (
	"os"
	"testing"
)

func TestArchiveLogs(t *testing.T) {
	// create a temporary log file
	logFile, err := os.CreateTemp("", "test-log-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp log file: %v", err)
	}
	defer os.Remove(logFile.Name())

	// Write some logs to the file
	logContent := "This is a test log line"
	if _, err := logFile.WriteString(logContent + "\n"); err != nil {
		t.Fatalf("Failed to write to log file: %v", err)
	}

	// Archive the logs
	archiveDir := "test-archive"
	archivePath, err := ArchiveLogs(logFile.Name(), archiveDir, false)
	if err != nil {
		t.Fatalf("Failed to archive logs: %v", err)
	}
	defer os.RemoveAll(archiveDir)

	// Verify the archive file exist
	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		t.Fatalf("Archive file was not created: %v", err)
	}
}
