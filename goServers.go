package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()

	channel := make(chan string)

	servers := []string{
		"http://google.com",
		"http://fb.com",
		"http://twitter.com",
	}

	checkServers(servers)

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	endTime := time.Since(startTime)

	fmt.Printf("time: %s", endTime)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println("Server ", server, " is DOWN.")
		channel <- "Server " + server + " is DOWN."
	} else {
		fmt.Println("Server ", server, " is UP.")
		channel <- "Server " + server + " is UP."
	}
}

func checkServers(servers []string, channel chan string) {
	for _, server := range servers {
		go checkServer(server, channel)
	}
}
