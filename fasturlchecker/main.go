package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("REQUEST FAILED")

type result struct {
	url    string
	status string
}

func (r result) String() string {
	return fmt.Sprint(r.url, ": ", r.status)
}

func main() {
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.naver.com",
		"https://www.soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
	}

	c := make(chan result)

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

// send-only channel
func hitUrl(url string, c chan<- result) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		c <- result{url: url, status: "FAILED"}
	} else {
		c <- result{url: url, status: "OK"}
	}
}
