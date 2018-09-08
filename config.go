package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Config holds configurable settings.
type Config struct {
	CaptureCommand     string `json:"capture_command"`
	CaptureCommandArgs string `json:"capture_command_args"`
	FilePath           string `json:"file_path"`
	DropboxToken       string `json:"dropbox_token"`
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
