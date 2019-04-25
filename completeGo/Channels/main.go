package main

import (
	"fmt"
	"net/http"
)

func main() {

	s := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://pivotal.io",
	}

	c := make(chan string)

	for i := range s {
		go checkLink(s[i], c)
	}

	//for j := 0; j < len(s); j++ {
	//	fmt.Println(<-c)
	//}

	for l := range c {
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		c <- "Might be down I think"
	}

	fmt.Println(link, "is up!")
	c <- "It's online!"
}

//A channel is the bridge||interceptor that provides communication across processes (go routines).
