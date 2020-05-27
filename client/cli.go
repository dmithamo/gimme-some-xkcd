// Package main defines the clients of the app
// In this file: a cli-based client
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dmithamo/gimme-some-xkcd/pkg/xkcdapi"
)

func main() {
	comicNumber := flag.Int("n", 0, "Number of the comic to retrieve")
	saveComic := flag.Bool("s", false, "Choose whether to save comic to disk <saved in ~/Downloads/xkcd directory> ")
	flag.Parse()

	api := xkcdapi.XKCD{Client: &http.Client{}}
	resp, err := api.Get(*comicNumber)

	if err != nil {
		fmt.Println("ERROR_FETCH", err)
	}

	comic := resp.FormatResponse()
	prettyPrinted, err := comic.Jsonify()
	if err != nil {
		fmt.Println("ERROR_JSONIFY", err)
	}
	fmt.Println(prettyPrinted)

	if *saveComic {
		err = saveToDisk(comic.Image, fmt.Sprintf("%v.png", comic.Number))
		if err != nil {
			fmt.Println("ERROR_SAVE", err)
		}
	}
}

func saveToDisk(imageURL, saveAs string) error {
	httpClient := http.Client{}
	resp, err := httpClient.Get(imageURL)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(saveAs, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
