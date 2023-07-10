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

// Check the connection to API URLs by simply
// verifying the response status of both URLs.
func CheckAPIURLConnection() error {
	resp, err := http.Head(fmt.Sprintf(API_URL_GET_USER_INFORMATION, "instagram"))
	if err != nil {
		return fmt.Errorf("error calling the API request: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error calling API endpoint '%s', code given: %d [%s]", resp.Request.URL, resp.StatusCode, resp.Status)
	}

	resp, err = http.Head(fmt.Sprintf(API_URL_GET_USER_INFORMATION, "instagram"))
	if err != nil {
		return fmt.Errorf("error calling the API request: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error calling API endpoint '%s', code given: %d [%s]", resp.Request.URL, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	return nil
}

// Call API request to get user information.
func GetUserInformation(username string) (*UserInformation, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL_GET_USER_INFORMATION, username))
	if err != nil {
		return nil, fmt.Errorf("error calling the API request: %s", err.Error())
	}
	defer resp.Body.Close()

	var userInformation UserInformation
	err = json.NewDecoder(resp.Body).Decode(&userInformation)
	if err != nil {
		return nil, fmt.Errorf("error decoding the response body: %s", err.Error())
	}

	return &userInformation, nil
}

// Call API request to get user stories.
func GetUserStories(userInformation *UserInformation) (*Story, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL_GET_STORY, userInformation.Result.User.Pk))
	if err != nil {
		return nil, fmt.Errorf("error calling the API request: %s", err.Error())
	}
	defer resp.Body.Close()

	var story Story
	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {
		return nil, fmt.Errorf("error decoding the response body: %s", err.Error())
	}

	return &story, nil
}

// Parse the stories by separating
// the photos or videos by content types.
func ParseStory(story *Story) *[]File {
	files := make([]File, len(story.Result))

	for i, r := range story.Result {
		newFile := File{
			Name: strconv.Itoa(int(time.Now().UnixNano())),
		}

		if len(r.VideoVersions) > 0 {
			newFile.Extension = ".mp4"
			newFile.URL = r.VideoVersions[0].URL

			files[i] = newFile
		} else {
			newFile.Extension = ".jpg"
			newFile.URL = r.ImageVersions2.Candidates[0].URL

			files[i] = newFile
		}

		// When running on Windows, the time.Now() function returns the same time precision.
		// This causes the file name to be the same for each file. To fix this issue, add a delay.
		// Reference: https://stackoverflow.com/questions/57285292/why-does-time-now-unixnano-returns-the-same-result-after-an-io-operation
		time.Sleep(time.Millisecond)
	}

	return &files
}
