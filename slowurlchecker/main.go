package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("REQUEST FAILED")

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
	}

	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)

		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
