package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Alphabets = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 5
	buf := make([]byte, n)
	for i, _ := range buf {
		buf[i] = Alphabets[rand.Intn(len(Alphabets))]
	}

	fmt.Println(string(buf))
}
