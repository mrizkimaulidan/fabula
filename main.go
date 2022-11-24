package main

import (
	"log"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/parser"
)

func main() {
	parser := parser.NewParser()

	resp, err := parser.Call()
	if err != nil {
		log.Fatal(err.Error())
	}

	fileS := file.NewFile()
	fileS.CreateDir()

	files := parser.Parsing(resp)

	for _, f := range *files {
		resp, err := fileS.GetFile(f.URL)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()

		createdFile, err := fileS.CreateFile(f, resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer createdFile.Close()
	}
}
