package instagram

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type InstagramInterface interface {
	GetInstagramProfile(username string) (*Instagram, error)
}

type Instagram struct {
	ProfileID string
	Username  string
}

func NewInstagram() InstagramInterface {
	return &Instagram{}
}

func (i *Instagram) GetInstagramProfile(username string) (*Instagram, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/%s", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return nil, errors.New("too many request, try again later")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	profile := Instagram{
		Username:  username,
		ProfileID: i.extractValue(string(responseBody), "profile_id"),
	}

	return &profile, nil
}

func (i *Instagram) extractValue(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.Split(match, ":")

	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}
