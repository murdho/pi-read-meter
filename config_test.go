package main

import (
	"regexp"
	"testing"
)

const (
	testConfigFile = "config.example.json"
)

var testConfig = Config{
	CaptureCommand:     "raspistill",
	CaptureCommandArgs: "-vf -hf -o %s",
	FilePath:           "~/pi-read-meter/image_%s.jpg",
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
