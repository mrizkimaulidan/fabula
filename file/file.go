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

type File struct {
	Filename  string
	Extension string
	URL       string
	Instagram *instagram.Instagram
}

func NewFile(instagram *instagram.Instagram) *File {
	return &File{
		Instagram: instagram,
	}
}

// Show download text on console
func (f *File) DownloadText(file *File) {
	log.Printf("downloading.. %s[%s]", file.Filename, file.Extension)
}

// Print output information where is the
// story has been downloaded
func (f *File) OutputPath() {
	log.Printf("stories saved on : %s/%s", DIR, f.Instagram.Username)
}

// Create directory for saving the story
// file
func (f *File) CreateDir() error {
	// stories/{instagram-username}
	path := fmt.Sprintf("%s/%s", DIR, f.Instagram.Username)
	return os.MkdirAll(path, os.ModePerm)
}

// Create file where is content are
// the story from source reader
// on parameter
func (f *File) CreateFile(file File, source io.Reader) (*os.File, error) {
	// stories/{instagram-username}/{unixTime}.{extension}
	fullPath := fmt.Sprintf("%s/%s/%s%s", DIR, f.Instagram.Username, file.Filename, file.Extension)
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

// Get file from URL provided on parameter
func (f *File) GetFile(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("error getting file %s", err.Error())
	}

	return resp, nil
}
