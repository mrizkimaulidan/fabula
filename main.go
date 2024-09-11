package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

var (
	username string
	option   string
)

func main() {
	flag.StringVar(&username, "username", "", "The Instagram username")
	flag.StringVar(&option, "option", "", "The parsing option 'story' or 'highlight'")
	flag.Parse()

	// handle if username is not provided in the flag
	if username == "" {
		log.Fatal("Username not provided in the flag. use --help flag for more info")
	}

	// handle if options is not provided in the flag
	if option == "" {
		log.Fatal("Option not provided in the flag. use --help flag for more info")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Checking connection to the API please wait..")
	err := CheckAPIURLConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("The connection seems ok!")

	userInformation, err := GetUserInformation(username)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch option {
	case "story":
		directoryName := "stories"
		userStories, err := GetUserStories(userInformation)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = CreateDir(fmt.Sprintf("%s/%s", directoryName, userInformation.Result.User.Username))
		if err != nil {
			log.Fatal(err.Error())
		}

		storyFiles := ParseStory(userStories)

		fmt.Println("=======================================")
		fmt.Printf("= Name\t\t: %s (@%s)\n", userInformation.Result.User.FullName, userInformation.Result.User.Username)
		fmt.Printf("= Followers\t: %d\n", userInformation.Result.User.FollowerCount)
		fmt.Printf("= Followings\t: %d\n", userInformation.Result.User.FollowingCount)
		fmt.Printf("= Public Email\t: %s\n", userInformation.Result.User.PublicEmail)
		fmt.Println("=======================================")
		fmt.Printf("Found %d stories for the user\n", len(*storyFiles))

		var wg sync.WaitGroup
		for _, f := range *storyFiles {
			wg.Add(1)
			go func(f File) {
				defer wg.Done()
				fmt.Printf("Downloading... %s%s\n", f.Name, f.Extension)

				fileStream, err := GetFile(f.URL)
				if err != nil {
					log.Fatal(err.Error())
				}
				defer fileStream.Body.Close()

				createdFileStream, err := CreateFile("stories", f, *userInformation, fileStream.Body)
				if err != nil {
					log.Fatal(err.Error())
				}
				defer createdFileStream.Close()

				fmt.Printf("Downloaded... %s%s\n", f.Name, f.Extension)
			}(f)
		}
		wg.Wait()

		fmt.Println("All stories have been downloaded!")

	case "highlight":
		directoryName := "highlights"
		userHightlightList, err := GetUserStoryHighlights(userInformation)
		if err != nil {
			log.Fatal(err.Error())
		}

		var highlights []struct {
			Number int
			ID     string
			Title  string
		}

		for i, v := range userHightlightList.Result {
			highlights = append(highlights, struct {
				Number int
				ID     string
				Title  string
			}{
				Number: i + 1,
				ID:     v.ID,
				Title:  v.Title,
			})
		}

		fmt.Println("=======================================")
		fmt.Printf("= Found the users with %d story highlights\n", len(highlights))
		for _, v := range highlights {
			fmt.Printf("= %d. %s\n", v.Number, v.Title)
		}
		fmt.Println("=======================================")

		var highlightNumber int
		fmt.Print("Which story highlights you want to download (in number): ")
		fmt.Scan(&highlightNumber)

		for _, v := range highlights {
			if v.Number == highlightNumber {
				userHighlightStories, err := GetUserHighlightStory(v.ID)
				if err != nil {
					log.Fatal(err.Error())
				}

				hightlightFiles := ParseHighlightStory(userHighlightStories)

				err = CreateDir(fmt.Sprintf("%s/%s", directoryName, userInformation.Result.User.Username))
				if err != nil {
					log.Fatal(err.Error())
				}

				var wg sync.WaitGroup
				for _, f := range *hightlightFiles {
					wg.Add(1)
					go func(f File) {
						defer wg.Done()
						fmt.Printf("Downloading... %s%s\n", f.Name, f.Extension)

						fileStream, err := GetFile(f.URL)
						if err != nil {
							log.Fatal(err.Error())
						}
						defer fileStream.Body.Close()

						createdFileStream, err := CreateFile("highlights", f, *userInformation, fileStream.Body)
						if err != nil {
							log.Fatal(err.Error())
						}
						defer createdFileStream.Close()

						fmt.Printf("Downloaded... %s%s\n", f.Name, f.Extension)
					}(f)
				}
				wg.Wait()

				fmt.Println("All highlight story have been downloaded!")
			}
		}
	default:
		log.Fatal("Option are not available! Please check --help for more information.")
	}
}
