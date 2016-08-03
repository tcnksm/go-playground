package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {

	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Millisecond)
	defer cancel()

	log.Printf("%#v", ctx.Value(http.ServerContextKey))
	log.Printf("%#v", ctx.Value(http.LocalAddrContextKey))

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// ...
}
