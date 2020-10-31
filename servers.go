package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()

	servers := []string{
		"http://google.com",
		"http://fb.com",
		"http://twitter.com",
	}

	checkServers(servers)

	endTime := time.Since(startTime)

	fmt.Printf("time: %s", endTime)
}

func checkServer(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println("Server ", server, " is DOWN.")
	} else {
		fmt.Println("Server ", server, " is UP.")
	}
}

func checkServers(servers []string) {
	for _, server := range servers {
		checkServer(server)
	}
}
