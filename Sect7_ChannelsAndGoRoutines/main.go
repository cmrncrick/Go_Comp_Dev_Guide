package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// Creating slice of strings
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Creating new channel
	c := make(chan string)

	// Iterating over slice of strings and performing checkLink
	for _, link := range links {
		go checkLink(link, c)
	}

	// Receiving value out of this channel and printing it
	for l := range c {
		// Declaring function literal (lambda) function 
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	// Checking for an error during Get request
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		
		// Into the channel, send this message
		c <- link
		return
	}

	// If no error, assume link is up
	fmt.Println(link, "is up!")

	// Into the channel, send this message
	c <- link
}