// +build !go1.7

package main

import "golang.org/x/net/context"

func Foo() {
	_ = context.TODO()
}
