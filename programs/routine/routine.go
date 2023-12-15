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

	for l := range c {
		func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	fmt.Println("current link : ", link) //When a link is received on the channel, you start a new goroutine (go checkLink(l, c)) to check the same link again after a 5-second sleep.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
