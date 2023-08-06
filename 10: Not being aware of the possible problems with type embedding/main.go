package main

import (
	"fmt"

	"example.com/test/bar"
)

type Foo struct {
	bar.Bar
}

func main() {
	foo := Foo{}
	foo.SetBaz(42)
	fmt.Println(foo.Baz())
}
