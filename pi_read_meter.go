package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Configurable:
// - image location & filename with timestamp
// - Dropbox credentials

// Execution interval is on caller (cron, etc)

// Steps:
// Load config
// Capture image
// Upload image to Dropbox

// Additional:
// Logging: start, steps (success or error), end

// ---

// Config holds configurable settings.
type Config struct {
	CaptureCommandPattern string `json:"capture_command_pattern"`
	ImageDir              string `json:"image_dir"`
	FilenamePattern       string `json:"filename_pattern"`
	DropboxToken          string `json:"dropbox_token"`
}

// LoadConfig loads config from configuration file to Config struct.
func LoadConfig(configFilename string) (*Config, error) {
	b, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file failed")
	}

	config := new(Config)
	if err := json.Unmarshal(b, config); err != nil {
		return nil, errors.Wrap(err, "unmarshalling json failed")
	}

	return config, nil
}
