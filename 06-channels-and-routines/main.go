package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		go (func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		})(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Something wrong with " + link)
		// c <- "Problem here"
		c <- link
		return
	}

	fmt.Println(link + " is up")
	c <- link
}
