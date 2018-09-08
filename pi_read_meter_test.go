package main

import "testing"

const (
	testConfigFile = "config.example.json"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig(testConfigFile)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expectedConfig := Config{
		CaptureCommandPattern: "raspistill -vf -hf -o %s",
		ImageDir:              "~/pi-read-meter/images",
		FilenamePattern:       "image_%s.jpg",
		DropboxToken:          "fcaea99c-67fa-4a56-9d08-6e133fd05759",
	}

	if config != nil && expectedConfig != *config {
		t.Errorf("\nexpected %#v, \ngot %#v", expectedConfig, config)
	}
}
