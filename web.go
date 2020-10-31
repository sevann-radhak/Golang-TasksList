package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWirter struct{}

func (webWirter) Write(p []byte) (int, error) {
	fmt.Println(string(p))

	return len(p), nil
}

func main() {
	result, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println(error)
	} else {
		e := webWirter{}
		io.Copy(e, result.Body)
	}
}
