package main

import (
	"testing"
)

const (
	testConfigFile = "config.example.json"
)

var testConfig = Config{
	CaptureCommand:     "raspistill",
	CaptureCommandArgs: "-vf -hf -o %s",
	FilePath:           "/home/pi/pi-read-meter/image_%s.jpg",
	DropboxToken:       "fcaea99c-67fa-4a56-9d08-6e133fd05759",
}

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig(testConfigFile)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if config != nil && testConfig != *config {
		t.Errorf("\nexpected %#v, \ngot %#v", testConfig, config)
	}
}
