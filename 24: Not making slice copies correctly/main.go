package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	fmt.Printf("src slice: %#+v\n", src)
	printResult := func(n string, s []int, c int) {
		fmt.Printf("%s: %#+v, %d copied of %d\n", n, s, c, len(src))
	}

	var dst1 []int
	c1 := copy(dst1, src)
	printResult("dst1", dst1, c1)

	dst2 := make([]int, len(src)-1)
	c2 := copy(dst2, src)
	printResult("dst1", dst2, c2)

	dst3 := make([]int, len(src))
	c3 := copy(dst3, src)
	printResult("dst3", dst3, c3)

	dst4 := append([]int(nil), src...)
	printResult("dst4", dst4, len(dst4))
}
