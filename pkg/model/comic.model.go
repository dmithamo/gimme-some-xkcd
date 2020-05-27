// Package model defines the structure of the application's data
package model

import (
	"encoding/json"
	"fmt"
)

// ComicResponse is the struct-to-json mapping of the response
type ComicResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// ComicModel is the struct-to-json mapping of the Comic
type ComicModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Image       string `json:"image"`
}

// FormatDate returns a formatted string of the date on the response
func (r ComicResponse) FormatDate() string {
	return fmt.Sprintf("%s-%s-%s", r.Year, r.Month, r.Day)
}

// FormatResponse extracts what we deem useful from the api's reponse
func (r ComicResponse) FormatResponse() ComicModel {
	return ComicModel{
		Title:       r.Title,
		Description: r.Alt,
		Image:       r.Img,
		Number:      r.Num,
		Date:        r.FormatDate(),
	}
}

// PrettyMeta formats meta data of the comic
func (c ComicModel) PrettyMeta() string {
	return fmt.Sprintf("%v | %s | %s | %s", c.Number, c.Date, c.Title, c.Description)
}

// Jsonify converts a Comic struct into json
func (c ComicModel) Jsonify() (string, error) {
	jsonified, err := json.Marshal(c)
	return string(jsonified), err
}
