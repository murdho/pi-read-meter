package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"time"

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

// IShellExec executes a shell command.
type IShellExec interface {
	Exec(commandAndArgs string) error
}

// ShellExec is a shell command executer.
type ShellExec struct{}

// Exec runs a shell command.
func (*ShellExec) Exec(commandAndArgs string) error {
	cmd := exec.Command(commandAndArgs)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "executing shell command failed")
	}

	return nil
}

// CaptureImage captures image by executing a shell command and returns a path to the image.
func CaptureImage(shellExec IShellExec, config *Config) (string, error) {
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	filename := fmt.Sprintf(config.FilenamePattern, timestamp)
	imagePath := path.Join(config.ImageDir, filename)
	command := fmt.Sprintf(config.CaptureCommandPattern, imagePath)

	shellExec.Exec(command)

	return imagePath, nil
}
