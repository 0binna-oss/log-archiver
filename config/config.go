package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ArchiveDir    string `json:"archiveDir"`
	Compress      bool   `json:"compress"`
	MaxLogSizeMB  int    `json:"maxLogSizeMB"`
	MaxLogAgeDays int    `json:"maxLgAgeDays"`
}

// LoadConfig loads the configuration from the json file
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	return &config, nil
}
