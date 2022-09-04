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

func (f *File) GetFileExtension(url string) string {
	s := strings.Split(url, "/")
	filename := s[len(s)-1]
	extension := strings.Split(filename, ".")[1]

	return extension
}

func (f *File) SetExtension(e string) {
	f.Extension = e
}

func (f *File) SetInstagramUsername(u string) {
	f.InstagramUsername = u
}

func (f *File) CreateDir() error {
	return os.MkdirAll(fmt.Sprintf("stories/%s", f.InstagramUsername), os.ModePerm)
}

func (f *File) CreateFile() (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("stories/%s/%d.%s", f.InstagramUsername, time.Now().UnixNano()/1000000, f.Extension))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) CopyFile(destination io.Writer, source io.Reader) error {
	_, err := io.Copy(destination, source)

	return err
}

func NewFile() *File {
	return &File{}
}
