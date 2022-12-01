package main

import (
	"flag"
	"log"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
	"github.com/mrizkimaulidan/fabula/parser"
)

func main() {
	var username string
	flag.StringVar(&username, "username", "", "the instagram username")
	flag.Parse()

	instagram := instagram.NewInstagram()
	instagramProfile, err := instagram.GetInstagramProfile(username)
	if err != nil {
		log.Fatal(err.Error())
	}

	parser := parser.NewParser(*instagramProfile)
	response, err := parser.Call()
	if err != nil {
		log.Fatal(err.Error())
	}

	files := parser.Parsing(response)
	file := file.NewFile(*instagramProfile)

	for _, f := range *files {
		response, err := file.GetFile(f.URL)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer response.Body.Close()

		createdFile, err := file.CreateFile(f, response.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer createdFile.Close()
	}
}
