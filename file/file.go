package file

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mrizkimaulidan/fabula/instagram"
)

const DIR = "stories"

type FileInterface interface {
	CreateDir() error
	CreateFile(file File, source io.Reader) (*os.File, error)
	GetFile(URL string) (*http.Response, error)
	OutputPath()
}

type File struct {
	Filename  string
	Extension string
	URL       string
	Instagram *instagram.Instagram
}

func NewFile(instagram *instagram.Instagram) FileInterface {
	return &File{
		Instagram: instagram,
	}
}

func (f *File) OutputPath() {
	log.Printf("stories saved on : %s/%s", DIR, f.Instagram.Username)
}

func (f *File) CreateDir() error {
	// stories/{instagram-username}
	path := fmt.Sprintf("%s/%s", DIR, f.Instagram.Username)
	return os.MkdirAll(path, os.ModePerm)
}

func (f *File) CreateFile(file File, source io.Reader) (*os.File, error) {
	// stories/{instagram-username}/{unixTime}.{extension}
	fullPath := fmt.Sprintf("%s/%s/%s%s", DIR, f.Instagram.Username, file.Filename, file.Extension)
	createdFile, err := os.Create(fullPath)
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
