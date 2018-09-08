package main

import (
	"regexp"
	"strings"
	"testing"
)

const (
	testConfigFile = "config.example.json"
)

var testConfig = Config{
	CaptureCommand:     "raspistill",
	CaptureCommandArgs: "-vf -hf -o %s",
	ImageDir:           "~/pi-read-meter/images",
	FilenamePattern:    "image_%s.jpg",
	DropboxToken:       "fcaea99c-67fa-4a56-9d08-6e133fd05759",
}

var filenameRegex = regexp.MustCompile(`image_\d{4}-(\d{2}-){4}\d{2}.jpg`)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig(testConfigFile)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if config != nil && testConfig != *config {
		t.Errorf("\nexpected %#v, \ngot %#v", testConfig, config)
	}
}

// make dir for images
// exec capture command
// check that file got created
// return path to file

type testShellExec struct {
	executedCommands []string
}

func (t *testShellExec) Exec(command string, args ...string) error {
	cmdAndArgs := append([]string{command}, args...)
	t.executedCommands = append(t.executedCommands, strings.Join(cmdAndArgs, " "))
	return nil
}

func TestCaptureImage(t *testing.T) {
	shellExec := new(testShellExec)
	imagePath, err := CaptureImage(shellExec, &testConfig)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(shellExec.executedCommands) != 1 {
		t.Errorf("expected executed commands '%s' to have 1 element", shellExec.executedCommands)
	}

	if !strings.HasPrefix(imagePath, testConfig.ImageDir) {
		t.Errorf("expected image path '%s' to have prefix '%s'", imagePath, testConfig.ImageDir)
	}

	if !filenameRegex.MatchString(imagePath) {
		t.Errorf("expected image path '%s' to match '%s'", imagePath, filenameRegex.String())
	}
}
