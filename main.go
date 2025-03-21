package main

import (
	"fmt"
	"log"
	"os"

	"log-archiver/archiver"
	"log-archiver/config"
	"log-archiver/rotator"
	"log-archiver/utils"
)

func main() {
	// load configuration
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	// Checks if the user provided a log file as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: log-archiver <logfile>")
		return
	}

	logFilePath := os.Args[1]

	// Validate the log file path
	if err := utils.ValidateFilePath(logFilePath); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Archive the logs
	archiveFilePath, err := archiver.ArchiveLogs(logFilePath, config.ArchiveDir, config.Compress)
	if err != nil {
		log.Fatalf("Error archiving logs: %v\n", err)
	}

	fmt.Printf("Logs archived successfully to %s\n", archiveFilePath)

	// Rotate logs
	if err := rotator.RotateLogs(config.ArchiveDir, config.MaxLogSizeMB, config.MaxLogAgeDays); err != nil {
		log.Fatalf("Error rotating logs: %v\n", err)
	}

	fmt.Println("Log rotation completed successfully")
}
