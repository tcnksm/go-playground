package main

import (
	"bytes"
	"fmt"

	"golang.org/x/net/http2"
)

func main() {
	buf := new(bytes.Buffer)
	fr := http2.NewFramer(buf, buf)

	var streamID uint32 = 1<<24 + 2<<16 + 3<<8 + 4
	fr.WriteData(streamID, true, []byte("Hello"))

	b := buf.Bytes()
	fmt.Printf("Frame: %q\n", b)

	fmt.Printf("Type: %x\n", b[4:5])
	fmt.Printf("StremID: %x\n", b[5:9])
	fmt.Printf("DATA: %x\n", b[9:])
}
