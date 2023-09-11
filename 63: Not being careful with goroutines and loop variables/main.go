package main

import (
	"fmt"
	"sync"
)

func badSum(s []int) int {
	wg := sync.WaitGroup{}
	wg.Add(len(s))

	var res int
	for _, v := range s {
		go func() {
			defer wg.Done()
			res += v
		}()
	}

	wg.Wait()
	return res
}

func okSum1(s []int) int {
	wg := sync.WaitGroup{}
	wg.Add(len(s))

	var res int
	for _, v := range s {
		go func(val int) {
			defer wg.Done()
			res += val
		}(v)
	}

	wg.Wait()
	return res
}

func okSum2(s []int) int {
	wg := sync.WaitGroup{}
	wg.Add(len(s))

	var res int
	for _, v := range s {
		val := v
		go func() {
			defer wg.Done()
			res += val
		}()
	}

	wg.Wait()
	return res
}

func main() {
	s := []int{1, 5, 10}

	for {
		fmt.Printf("badSum(): %d\n", badSum(s))
		fmt.Printf("okSum1(): %d\n", okSum1(s))
		fmt.Printf("okSum2(): %d\n", okSum2(s))
		fmt.Scanln()
	}
}
