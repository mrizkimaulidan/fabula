package file

import (
	"io"
	"net/http"
	"os"
)

const DIR = "stories"

type FileInterface interface {
	CreateDir() error
	CreateFile(file File, source io.Reader) (*os.File, error)
	GetFile(URL string) (*http.Response, error)
}

type File struct {
	Filename  string
	Extension string
	URL       string
}

func NewFile() FileInterface {
	return &File{}
}

func (f *File) CreateDir() error {
	return os.MkdirAll(DIR, os.ModePerm)
}

func (f *File) CreateFile(file File, source io.Reader) (*os.File, error) {
	createdFile, err := os.Create(DIR + "/" + file.Filename + file.Extension)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(createdFile, source)
	if err != nil {
		return nil, err
	}

	return createdFile, nil
}

func (f *File) GetFile(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
