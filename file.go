package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type File struct {
	Name      string
	URL       string
	Extension string
}

// Create a directory folder based
// on the given name.
func CreateDir(name string) error {
	path := path.Join(name)

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
func CreateFile(dir string, file File, source io.Reader) (*os.File, error) {
	fullPath := path.Join(dir, file.Name+file.Extension)
	createdFile, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error creating file: %s", err.Error())
	}

	bufferWriter := bufio.NewWriter(createdFile)
	defer bufferWriter.Flush()

	_, err = io.Copy(bufferWriter, source)
	if err != nil {
		return nil, fmt.Errorf("error copying from source: %s", err.Error())
	}

	return createdFile, nil
}
