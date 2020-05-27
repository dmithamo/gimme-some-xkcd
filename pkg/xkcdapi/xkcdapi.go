// Package xkcdapi provides methods for interfacing with the xkcd api
package xkcdapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	cm "github.com/dmithamo/gimme-some-xkcd/pkg/comicmodel"
)

const BaseURL string = "https://xkcd.com"

// XKCD structures the interface to the xkcd api
type XKCD struct {
	Client *http.Client
}

// Get fetches a comic, given comic number
func (x XKCD) Get(comicNumber int) (cm.ComicResponse, error) {
	resp, err := x.Client.Get(x.buildURL(comicNumber))

	if err != nil {
		return cm.ComicResponse{}, err
	}
	defer resp.Body.Close()

	var retrievedComic cm.ComicResponse
	if err := json.NewDecoder(resp.Body).Decode(&retrievedComic); err != nil {
		return cm.ComicResponse{}, err
	}

	return retrievedComic, nil
}

// buildURL creates an explicit url to a comic
func (x XKCD) buildURL(comicNumber int) string {
	// latest comic url
	if comicNumber == 0 {
		return fmt.Sprintf("%s/info.0.json", BaseURL)
	}

	// specific comic url
	return fmt.Sprintf("%s/%v/info.0.json", BaseURL, comicNumber)
}
