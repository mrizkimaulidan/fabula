package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	API_URL_GET_USER_INFORMATION = "https://storiesig.info/api/ig/userInfoByUsername/%s"
	API_URL_GET_STORY            = "https://storiesig.info/api/ig/stories/%s"
)

func GetUserInformation(username string) (*UserInformation, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL_GET_USER_INFORMATION, username))
	if err != nil {
		return nil, fmt.Errorf("error cok %s", err.Error())
	}
	defer resp.Body.Close()

	var userInformation UserInformation
	err = json.NewDecoder(resp.Body).Decode(&userInformation)
	if err != nil {
		return nil, fmt.Errorf("error cok %s", err.Error())
	}

	return &userInformation, nil
}

func GetUserStories(userInformation *UserInformation) (*Story, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL_GET_STORY, userInformation.Result.User.Pk))
	if err != nil {
		return nil, fmt.Errorf("error cok %s", err.Error())
	}
	defer resp.Body.Close()

	var story Story
	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {
		return nil, fmt.Errorf("error cok %s", err.Error())
	}

	return &story, nil
}

func ParsingStory(story *Story) *[]File {
	var (
		files = make([]File, len(story.Result))
	)

	for i, r := range story.Result {
		newFile := File{
			Name: strconv.Itoa(int(time.Now().UnixNano())),
		}

		if r.HasAudio {
			newFile.Extension = ".mp4"
			newFile.URL = r.VideoVersions[0].URL

			files[i] = newFile
		} else {
			newFile.Extension = ".jpg"
			newFile.URL = r.ImageVersions2.Candidates[0].URL

			files[i] = newFile
		}
	}

	return &files
}
