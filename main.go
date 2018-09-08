package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/tj/go-dropbox"
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

	filenameFull := fmt.Sprintf(config.FilePath, time.Now().Format("2006-01-02-15-04-05"))

	log.Printf("capturing -> %s", filenameFull)

	camera := NewCamera(config.CaptureCommand, config.CaptureCommandArgs)
	if err := camera.Capture(filenameFull); err != nil {
		log.Fatalf("capturing image failed: %s", err)
	}

	dropboxUploadFunc := func(name string, file io.Reader) error {
		dbx := dropbox.NewFiles(dropbox.NewConfig(config.DropboxToken))

		uploadInput := &dropbox.UploadInput{
			Path:       fmt.Sprintf("/%s", name),
			Mode:       dropbox.WriteModeAdd,
			AutoRename: false,
			Mute:       true,
			Reader:     file,
		}

		if _, err := dbx.Upload(uploadInput); err != nil {
			return errors.Wrap(err, "uploading file to Dropbox failed")
		}

		return nil
	}

	log.Printf("uploading <- %s", filenameFull)

	dropboxUploader := NewDropboxUploader(dropboxUploadFunc)
	if err := dropboxUploader.Upload(filenameFull); err != nil {
		log.Fatalf("uploading to Dropbox failed: %s", err)
	}

	log.Printf("done!        %s", filenameFull)
}
