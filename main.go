package main

import (
	"flag"
	"log"
	"sync"
)

var (
	username string
)

func main() {
	flag.StringVar(&username, "username", "", "the Instagram username")
	flag.Parse()

	// handle if username is not provided in the flag
	if username == "" {
		log.Fatal("username not provided in the flag. use --help flag for more info")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("checking connection to the API please wait..")
	err := CheckAPIURLConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("the connection seems ok!")

	userInformation, err := GetUserInformation(username)
	if err != nil {
		log.Fatal(err.Error())
	}

	userStories, err := GetUserStories(userInformation)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = CreateDir(userInformation.Result.User.Username)
	if err != nil {
		log.Fatal(err.Error())
	}

	storyFiles := ParseStory(userStories)

	log.Println("=======================================")
	log.Printf("Name\t\t: %s (@%s)", userInformation.Result.User.FullName, userInformation.Result.User.Username)
	log.Printf("Followers\t: %d", userInformation.Result.User.FollowerCount)
	log.Printf("Followings\t: %d", userInformation.Result.User.FollowingCount)
	log.Printf("Public Email\t: %s", userInformation.Result.User.PublicEmail)
	log.Println("=======================================")
	log.Printf("Found %d stories for the user", len(*storyFiles))

	var wg sync.WaitGroup
	for _, f := range *storyFiles {
		wg.Add(1)
		go func(f File) {
			defer wg.Done()
			log.Printf("Downloading... %s%s", f.Name, f.Extension)

			fileStream, err := GetFile(f.URL)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer fileStream.Body.Close()

			createdFileStream, err := CreateFile(f, *userInformation, fileStream.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer createdFileStream.Close()

			log.Printf("Downloaded... %s%s", f.Name, f.Extension)
		}(f)
	}
	wg.Wait()

	log.Println("All stories have been downloaded!")
}
