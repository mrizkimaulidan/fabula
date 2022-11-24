package main

import (
	"flag"
	"log"

	"github.com/mrizkimaulidan/fabula/instagram"
	"github.com/mrizkimaulidan/fabula/parser"
)

func main() {
	var username string
	flag.StringVar(&username, "username", "", "the instagram username")
	flag.Parse()

	instagram := instagram.NewInstagram()
	instagramProfile, err := instagram.GetProfileIDByUsername(username)
	if err != nil {
		log.Fatal(err.Error())
	}

	parser := parser.NewParser(*instagramProfile)
	parser.Start()
}
