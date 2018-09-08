package main

import (
	"fmt"

	"github.com/tj/go-dropbox"
)

// IDropboxUploader is an interface for a Dropbox file uploader.
type IDropboxUploader interface {
	Upload(localPath string) error
}

// DropboxUploader is an IDropboxUploader implementation.
type DropboxUploader struct {
	filesClient *dropbox.Files
}

// Upload uploads a local file to app directory in Dropbox.
func (du *DropboxUploader) Upload(localPath string) error {
	// TODO
	fmt.Printf("\033[1;33m%#v\033[0m\n", "Upload is not implemented yet")
	return nil
}

// NewDropboxUploader initializes a new DropboxUploader.
func NewDropboxUploader(token string) IDropboxUploader {
	return &DropboxUploader{
		filesClient: dropbox.NewFiles(dropbox.NewConfig(token)),
	}
}
