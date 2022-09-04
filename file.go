package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type File struct {
	InstagramUsername string
	Extension         string
}

// Getting the file extension on the story URL.
// Split the url to remove the `/` first.
// Get the filename, and then split the filename by dot
// separator. The last index is the extension.
func (f *File) GetFileExtension(url string) string {
	s := strings.Split(url, "/")
	filename := s[len(s)-1]
	extension := strings.Split(filename, ".")[1]

	return extension
}

// Set the file extension.
func (f *File) SetExtension(e string) {
	f.Extension = e
}

// Set the instagram username.
func (f *File) SetInstagramUsername(u string) {
	f.InstagramUsername = u
}

// Create directory folder for saving the story.
// The folder created in `stories/` and inside there
// will be another folder, the folder name is your
// target instagram username.
func (f *File) CreateDir() error {
	return os.MkdirAll(fmt.Sprintf("stories/%s", f.InstagramUsername), os.ModePerm)
}

// Create file of a story that has been downloaded.
// The file saved inside the `stories/username` folder.
func (f *File) CreateFile() (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("stories/%s/%d.%s", f.InstagramUsername, time.Now().UnixNano()/1000000, f.Extension))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Copying file to folder destination.
func (f *File) CopyFile(destination io.Writer, source io.Reader) error {
	_, err := io.Copy(destination, source)

	return err
}

func NewFile() *File {
	return &File{}
}
