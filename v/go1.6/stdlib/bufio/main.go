package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	l := 64 * 1024
	str := "a"
	for i := 0; i < l+10; i++ {
		str += "a"
	}
	br := bytes.NewReader([]byte(str + "\n"))
	scanner := bufio.NewScanner(br)
	scanner.Buffer([]byte{}, 2*l)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
