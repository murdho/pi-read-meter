package main

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// ICamera is an interface for an image capturing tool.
type ICamera interface {
	Capture(filePath string) error
}

// Camera is an image capturing tool.
type Camera struct {
	captureCommand string
	captureArgs    string
}

// Capture takes a pic to the specified file location.
func (c *Camera) Capture(filePath string) error {
	argsWithFile := fmt.Sprintf(c.captureArgs, filePath)
	if err := exec.Command(c.captureCommand, argsWithFile).Run(); err != nil {
		return errors.Wrap(err, "executing capture command failed")
	}

	return nil
}

// NewCamera initializes a camera ready to take pixxxxx.
func NewCamera(captureCommand, captureArgs string) ICamera {
	return &Camera{
		captureCommand: captureCommand,
		captureArgs:    captureArgs,
	}
}
