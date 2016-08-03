package foo

import (
	"fmt"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	cases := []struct {
		Base int
	}{
		{Base: 1},
		{Base: 10},
		{Base: 100},
	}

	for _, bc := range cases {
		b.Run(fmt.Sprintf("%d", bc.Base), func(b *testing.B) { benchFoo(b, bc.Base) })
	}
}

func BenchmarkFoo1(b *testing.B)   { benchFoo(b, 1) }
func BenchmarkFoo10(b *testing.B)  { benchFoo(b, 10) }
func BenchmarkFoo100(b *testing.B) { benchFoo(b, 100) }

func benchFoo(b *testing.B, base int) {
	for i := 0; i < b.N; i++ {
		Foo(base)
	}
}
