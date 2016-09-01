package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	"github.com/tcnksm/go-httptraceutil"
)

func main() {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	trace := httptrace.ClientTrace{
		GetConn: func(h string) {
			log.Println(h)
		},
	}

	ctx := httptraceutil.WithClientTrace(context.Background())
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	_ = res
}
