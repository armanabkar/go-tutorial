package main

import (
	"fmt"
	"net/http"
	"time"
)

// When we create a program, one 'Go routine' is automatically created (Main Routine), when there is a blocking call (like http request), it will wait
// By default Go use only one CPU core and Go Scheduler runs one thread on each 'logical' core
// Concurency is not Parallelism! Concurency means multiple threads executing code but Parallelism means multiple threads executed at he exact same time and requires multiple CPUs

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	// use Channels to communicate between routines
	channel := make(chan string) // create Channel

	for _, link := range links {
		// create new routine with 'go'
		// when Main Routine is completed, all Child Routines will be stopped
		go checkLink(link, channel)
	}

	// infinite loop
	for link := range channel {
		// Function Literal (Anonymous funcs)
		go func(link string) {
			time.Sleep(10 * time.Second) // 'time' package; to sleep routine

			checkLink(link, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down...")
		channel <- link // sending data to channel
		return
	}

	fmt.Println(link, "is up")
	channel <- link // sending data to channel
}
