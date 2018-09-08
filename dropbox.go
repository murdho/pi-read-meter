package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// NewDropboxUploader initializes a DropboxUploader.
func NewDropboxUploader(uploadFunc UploadFunc) *DropboxUploader {
	return &DropboxUploader{
		uploadFunc: uploadFunc,
	}
}

// DropboxUploader uploads a local file to Dropbox using custom uploading function.
type DropboxUploader struct {
	uploadFunc func(name string, contents io.Reader) error
}

// Upload uploads file from local path to Dropbox using specified function.
func (du *DropboxUploader) Upload(localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return errors.Wrap(err, "opening file for uploading failed")
	}

	filename := filepath.Base(localPath)
	if err := du.uploadFunc(filename, file); err != nil {
		return errors.Wrap(err, "uploading the file failed")
	}

	return nil
}

// UploadFunc is used to specify uploading logic in DropboxUploader.
type UploadFunc func(name string, contents io.Reader) error
