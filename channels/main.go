package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {

	flag.Bool

	links := []string{
		"http://facebook.com",
		"http://gmail.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)
	

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //calling go routine
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " is down")
		c <- link
		return

	}

	fmt.Println(link, " is up")
	c <- link

}
