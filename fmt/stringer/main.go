package main

import "fmt"

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/http context value " + k.name }

func main() {
	key := &contextKey{"http-server"}
	fmt.Println(key)
}
