package main

import (
	"time"
)

func main() {
	//Add websites here
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
		"http://apple.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}
