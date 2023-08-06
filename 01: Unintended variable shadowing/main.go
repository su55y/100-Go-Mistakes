package main

import "fmt"

func main() {
	var foo int
	if true {
		foo := 1
		fmt.Printf("foo inside if scope: %d\n", foo)
	} else {
		foo := 2
		fmt.Printf("foo inside if scope: %d\n", foo)
	}
	fmt.Printf("foo in main scope: %d\n", foo)
}
