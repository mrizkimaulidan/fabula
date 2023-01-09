package instagram

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const INSTAGRAM_URL = "https://www.instagram.com/%s"

type Instagram struct {
	ProfileID string
	Username  string
}

func New(username string) *Instagram {
	return &Instagram{
		Username: username,
	}
}

// Get the instagram profile from Instagram website
// we inspecting the HTML return response
// and fill up the ProfileID based on the
// instagram username
func (i *Instagram) GetInstagramProfile() (*Instagram, error) {
	resp, err := http.Get(fmt.Sprintf(INSTAGRAM_URL, i.Username))
	if err != nil {
		return nil, fmt.Errorf("error calling request to instagram %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return nil, errors.New("too many request, try again later")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body %s", err.Error())
	}

	i.ProfileID = i.extractValue(string(responseBody), "profile_id")

	return i, nil
}

// Extracting the HTML body response
// to be honest this is from stackoverflow lol
// i do not remember the original post
func (i *Instagram) extractValue(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.Split(match, ":")

	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}
