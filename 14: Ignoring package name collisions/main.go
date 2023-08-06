package main

import (
	"example.com/test/bar"
)

type Foo struct {
	bar.Bar
}

func main() {
	bar := bar.Bar{}
	bar2 := bar.Bar{}
}
