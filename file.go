package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const DIR = "stories"

type File struct {
	Name      string
	URL       string
	Extension string
}

// Create a directory folder based
// on the given name.
func CreateDir(name string) error {
	path := fmt.Sprintf("%s/%s", DIR, name)

	return os.MkdirAll(path, os.ModePerm)
}

// Get a file from a URL.
func GetFile(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Create a file based on the provided file struct,
// user information, and source reader, and save it into a folder.
func CreateFile(file File, userInformation UserInformation, source io.Reader) (*os.File, error) {
	// stories/{instagram-username}/{unixTime}.{extension}
	fullPath := path.Join(DIR, userInformation.Result.User.Username, file.Name+file.Extension)
	createdFile, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error creating file: %s", err.Error())
	}

	_, err = io.Copy(createdFile, source)
	if err != nil {
		return nil, fmt.Errorf("error copying from source: %s", err.Error())
	}

	return createdFile, nil
}
