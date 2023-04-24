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

	// for { // infinite loop
	// 	go checkLink(<-c, c)
	// }
	// line 25 is same as line 29
	for l := range c {
		// time.Sleep(2 * time.Second)
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up !")
	c <- link
}
