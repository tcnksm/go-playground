package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	l := 5
	str := "a"
	for i := 0; i < l; i++ {
		str += "a"
	}
	br := bytes.NewReader([]byte(str + "\n" + str + "\n"))
	scanner := bufio.NewScanner(br)
	scanner.Buffer([]byte("bbbbbbbbbbbbbbb\n"), 1)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
