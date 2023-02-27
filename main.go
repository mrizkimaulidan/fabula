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
	flag.StringVar(&username, "username", "", "the instagram username")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	storyFiles := ParsingStory(userStories)

	log.Println("=======================================")
	log.Printf("= Name\t\t: %s(@%s)", userInformation.Result.User.FullName, userInformation.Result.User.Username)
	log.Printf("= Followers\t: %d", userInformation.Result.User.FollowerCount)
	log.Printf("= Followings\t: %d", userInformation.Result.User.FollowingCount)
	log.Printf("= Public Email\t: %s", userInformation.Result.User.PublicEmail)
	log.Println("=======================================")
	log.Printf("Found the user with %d stories", len(*storyFiles))

	var wg sync.WaitGroup
	for _, f := range *storyFiles {
		wg.Add(1)
		go func(f File) {
			defer wg.Done()
			log.Printf("Downloading.. %s%s", f.Name, f.Extension)

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

			log.Printf("Downloaded.. %s%s", f.Name, f.Extension)
		}(f)
	}
	wg.Wait()

	log.Println("All stories has been downloaded!")
}
