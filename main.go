package main

import "fmt"

func main() {
	config := &Config{
		CaptureCommand:     "imagesnap",
		CaptureCommandArgs: "%s",
		FilenamePattern:    "image_%s.jpg",
		ImageDir:           "/Users/murdho/tmp/pi-images",
	}

	fmt.Println(CaptureImage(new(ShellExec), config))
}
