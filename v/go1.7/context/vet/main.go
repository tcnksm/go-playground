package main

import (
	"context"
	"time"
)

func main() {
	context, cancel := context.WithCancel()
	if false {
		cancel()
	} else {
		for {
			print(0)
			time.Sleep(1 * time.Second)
		}
	}
}
