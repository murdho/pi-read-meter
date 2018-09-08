package main

import (
	"io"
	"io/ioutil"
	"regexp"
	"testing"
)

var filenameRegex = regexp.MustCompile(`^image_\d{4}-(\d{2}-){4}\d{2}.jpg$`)

func TestDropboxUpload(t *testing.T) {
	var uploadFuncCalled bool

	uploadFunc := func(name string, file io.Reader) error {
		uploadFuncCalled = true

		if !filenameRegex.MatchString(name) {
			t.Errorf("expected filename '%s' to match pattern %s", name, filenameRegex.String())
		}

		b, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		expectedContents := "behind the image\n"
		if string(b) != expectedContents {
			t.Errorf("expected image file contents to be %#v, got %#v", expectedContents, string(b))
		}

		return nil
	}

	if err := NewDropboxUploader(uploadFunc).Upload("./testdata/image_2018-09-08-20-45-38.jpg"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if !uploadFuncCalled {
		t.Errorf("expected uploadFunc to be called")
	}
}
