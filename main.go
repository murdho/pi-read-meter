package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide config file as first argument")
	}

	configFile := os.Args[1]
	config, err := LoadConfig(configFile)
	if err != nil {
		log.Fatalf("loading config from '%s' failed: %s", configFile, err)
	}

	filename := fmt.Sprintf(config.FilePath, time.Now().Format("2006-01-02-15-04-05"))

	camera := NewCamera(config.CaptureCommand, config.CaptureCommandArgs)
	if err := camera.Capture(filename); err != nil {
		log.Fatalf("capturing image failed: %s", err)
	}

	dropboxUploader := NewDropboxUploader(config.DropboxToken)

	if err := dropboxUploader.Upload(filename); err != nil {
		log.Fatalf("uploading to Dropbox failed: %s", err)
	}
}
