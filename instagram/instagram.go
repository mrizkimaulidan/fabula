package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type InstagramInterface interface {
	GetProfileIDByUsername(username string) (*InstagramProfile, error)
}

type Instagram struct {
	ProfileID string
}

func NewInstagram() InstagramInterface {
	return &Instagram{}
}

func (i *Instagram) GetProfileIDByUsername(username string) (*InstagramProfile, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/web/search/topsearch/?query=%s", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(responseBody), "fail") {
		return nil, errors.New("rate limit reached, try again later")
	}

	var profile InstagramProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
