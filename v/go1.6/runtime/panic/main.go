package main

import "time"

func main() {
	go func() {
		time.Sleep(10 * time.Second)
	}()

	panic("yay")
}
