package main

import (
	"fmt"
	"io"
)

type V struct {
	x, y int
}

func (v *V) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, "v")
			return
		}

		fmt.Fprintf(s, "fff")
	case 's':
		io.WriteString(s, "string")
	}
}

func main() {
	v := &V{
		x: 1,
		y: 2,
	}

	fmt.Printf("%v\n", v)
	fmt.Printf("%+v\n", v)
	fmt.Printf("%s\n", v)
}
