package main

import (
	"bytes"
	"fmt"

	"golang.org/x/net/http2/hpack"
)

func main() {
	s := "www.example.com"

	fmt.Println(len(s))
	fmt.Println(hpack.HuffmanEncodeLength(s))

	b := hpack.AppendHuffmanString(nil, s)
	fmt.Printf("%x\n", b)

	var buf bytes.Buffer
	if _, err := hpack.HuffmanDecode(&buf, b); err != nil {
		fmt.Printf("[ERRRO] %s", err)
	}

	fmt.Printf("%s\n", buf.String())
}
