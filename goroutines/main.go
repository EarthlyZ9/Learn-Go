package main

import (
	"fmt"
	"time"
)

func main() {

	// this takes 20 seconds total
	// helloCount("jisoo")
	// helloCount("earthlyz9")

	// this takes 10 seconds total (concurrently)
	// go routine is alive while main function is alive
	// if main function finishes the go routines are killed
	// go helloCount("jisoo")
	// helloCount("earthlyz9")

	// hence this code does not work
	// go helloCount("jisoo")
	// go helloCount("earthlyz9")

	// make a channel to send the result of the helloCount to main
	c := make(chan string)
	people := [5]string{"jisoo", "earthlyz9", "leejisoo", "js", "ljs"}
	for _, person := range people {
		go helloCountWithChannel(person, c)
	}

	// receiving messages
	fmt.Println("Waiting for message")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // blocking operation (waiting for a message to be received)
	}
}

func helloCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello to", person)
		time.Sleep(time.Second * 1)
	}
}

func helloCountWithChannel(person string, channel chan string) {
	time.Sleep(time.Second * 5)

	// sending the result to main func
	channel <- "Hello to " + person
}
