package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://http2.golang.org/reqinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Printf("is HTTP2: %v (%s)\n", response.ProtoAtLeast(2, 0), response.Proto)
}
