package main

import (
	"flag"
	"log"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
	"github.com/mrizkimaulidan/fabula/parser"
)

func main() {
	var profileID string
	flag.StringVar(&profileID, "profileID", "", "the instagram profileID")
	flag.Parse()

	instagram := instagram.SetProfileID(profileID)
	parser := parser.NewParser(instagram.ProfileID)

	resp, err := parser.Call()
	if err != nil {
		log.Fatal(err.Error())
	}

	fileS := file.NewFile()
	fileS.CreateDir()

	files := parser.Parsing(resp)
	for _, f := range *files {
		// TODO: goroutine??
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
