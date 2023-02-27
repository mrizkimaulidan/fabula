package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const DIR = "stories"

type File struct {
	Name      string
	URL       string
	Extension string
}

// Create directory folder based
// on name at parameter.
func CreateDir(name string) error {
	path := fmt.Sprintf("%s/%s", DIR, name)

	return os.MkdirAll(path, os.ModePerm)
}

// Getting file from URL.
func GetFile(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Creating file based on source reader and saved it to
// folder.
func CreateFile(file File, userInformation UserInformation, source io.Reader) (*os.File, error) {
	// stories/{instagram-username}/{unixTime}.{extension}
	fullPath := fmt.Sprintf("%s/%s/%s%s", DIR, userInformation.Result.User.Username, file.Name, file.Extension)
	createdFile, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error creating file %s", err.Error())
	}

	_, err = io.Copy(createdFile, source)
	if err != nil {
		return nil, fmt.Errorf("error copying from source %s", err.Error())
	}

	return createdFile, nil
}
