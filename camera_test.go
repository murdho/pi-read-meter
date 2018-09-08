package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCameraCapture(t *testing.T) {
	os.Mkdir("tmp", os.ModePerm)

	testFile := fmt.Sprintf("./tmp/testfile_%d.txt", time.Now().Unix())
	defer os.Remove(testFile)

	if _, err := os.Stat(testFile); os.IsExist(err) {
		if err := os.Remove(testFile); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}

	camera := NewCamera("touch", "%s")
	if err := camera.Capture(testFile); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Errorf("expected file '%s' to exist", testFile)
	}
}
