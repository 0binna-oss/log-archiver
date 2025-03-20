package main

import (
	"fmt"
	"log"
	"os"

	"log-archiver/archiver"
	"log-archiver/config"
	"log-archiver/utils"
)

func main() {
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
	archiveFilePath, err := archiver.ArchiveLogs(logFilePath, config.DefaultArchiveDir)
	if err != nil {
		log.Fatalf("Error archiving logs: %v\n", err)
	}

	fmt.Printf("Logs archived successfully to %s\n", archiveFilePath)
}
