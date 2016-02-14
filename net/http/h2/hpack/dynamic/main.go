package main

import (
	"bytes"
	"fmt"

	"golang.org/x/net/http2/hpack"
)

func main() {
	var buf bytes.Buffer

	// First time
	e := hpack.NewEncoder(&buf)
	if err := e.WriteField(hpack.HeaderField{
		Name:  ":authority",
		Value: "www.example.com",
	}); err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("Encoded: %x (%d) \n", buf.Bytes(), len(buf.Bytes()))

	buf.Reset()

	// Second time
	if err := e.WriteField(hpack.HeaderField{
		Name:  ":authority",
		Value: "www.example.com",
	}); err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("Encoded: %x (%d) \n", buf.Bytes(), len(buf.Bytes()))
}
