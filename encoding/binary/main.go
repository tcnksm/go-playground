package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println(binary.BigEndian.Uint16([]byte("test")))
	fmt.Println(binary.BigEndian.Uint32([]byte("test")))
}
