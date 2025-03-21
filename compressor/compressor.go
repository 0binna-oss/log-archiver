package compressor

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// CompressFile compresses a file using gzip and returns the path to the compressed file
func CompressFile(filePath string) (string, error) {
	// Open the file to be compressed
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create the compressed file
	compressedFilePath := filePath + ".gz"
	compressedFile, err := os.Create(compressedFilePath)
	if err != nil {
		return "", fmt.Errorf("error creating compressed file: %w", err)
	}
	defer compressedFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(compressedFile)
	defer gzipWriter.Close()

	// Copy the file content to the gzip writer
	if _, err := io.Copy(gzipWriter, file); err != nil {
		return "", fmt.Errorf("error compressing file: %w", err)
	}

	return compressedFilePath, nil
}
