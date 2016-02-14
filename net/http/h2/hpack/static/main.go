package main

import (
	"bytes"
	"fmt"

	"golang.org/x/net/http2/hpack"
)

func main() {
	var buf bytes.Buffer

	e := hpack.NewEncoder(&buf)
	err := e.WriteField(hpack.HeaderField{
		Name:  ":method",
		Value: "GET",
	})

	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("Encoded: %x (%d) \n", buf.Bytes(), len(buf.Bytes()))

	var decoded hpack.HeaderField
	d := hpack.NewDecoder(4<<10, func(f hpack.HeaderField) {
		decoded = f
	})

	if _, err := d.Write(buf.Bytes()); err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("Decoded: %#v\n", decoded)
}
