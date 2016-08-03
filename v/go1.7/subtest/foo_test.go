package foo

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	for i := 1; i < 4; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) { testFoo(i, t) })
	}
}

func TestFoo1(t *testing.T) {
	testFoo(1, t)
}

func TestFoo2(t *testing.T) {
	testFoo(2, t)
}

func TestFoo3(t *testing.T) {
	testFoo(3, t)
}

func testFoo(i int, t *testing.T) {
	out, err := Foo(i)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	//..
	_ = out
}
