package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", "github.com:80")
	if err != nil {
		log.Fatal(err)
	}

	_ = conn
}
