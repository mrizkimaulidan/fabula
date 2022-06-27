package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const DOWNLOADED_FOLDER = "./stories"

type File struct {
	SubFolderName string
}

func NewFile() *File {
	return &File{}
}

// Change the SubFolder name
func (f *File) ChangeSubFolderName(s string) {
	f.SubFolderName = s
}

// Create directory to the filesystem
// if dir exists it will do nothing
// if not exists it will creating the folder
func (f *File) Mkdir() error {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s", DOWNLOADED_FOLDER, f.SubFolderName), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Creating file, this is a represent of
// images or videos of the instagram story
func (f *File) CreateFile(extension string) (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("%s/%s/%d.%s", DOWNLOADED_FOLDER, f.SubFolderName, time.Now().UnixNano(), extension))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Copying file from reader to destination
func (f *File) CopyFile(destination io.Writer, source io.Reader) error {
	_, err := io.Copy(destination, source)

	return err
}
