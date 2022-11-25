package file

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mrizkimaulidan/fabula/instagram"
)

const DIR = "stories"

type FileInterface interface {
	CreateDir() error
	CreateFile(file File, source io.Reader) (*os.File, error)
	GetFile(URL string) (*http.Response, error)
}

type File struct {
	Filename         string
	Extension        string
	URL              string
	InstagramProfile instagram.InstagramProfile
}

func NewFile(instagram instagram.InstagramProfile) FileInterface {
	return &File{
		InstagramProfile: instagram,
	}
}

func (f *File) CreateDir() error {
	// stories/{instagram-username}
	path := fmt.Sprintf("%s/%s", DIR, f.InstagramProfile.Users[0].User.Username)
	return os.MkdirAll(path, os.ModePerm)
}

func (f *File) CreateFile(file File, source io.Reader) (*os.File, error) {
	// stories/{instagram-username}/
	fullPath := fmt.Sprintf("%s/%s/%s%s", DIR, f.InstagramProfile.Users[0].User.Username, file.Filename, file.Extension)
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
