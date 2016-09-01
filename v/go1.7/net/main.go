package main

import (
	"log"
	"net/http"

	"github.com/tcnksm/go-httptraceutils"
)

func main() {

	url := "https://google.co.jp"
	req, _ := http.NewRequest("GET", url, nil)

	ctx := httptraceutils.WithClientTrace(req.Context())
	req = req.WithContext(ctx)

	client := http.DefaultClient
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
