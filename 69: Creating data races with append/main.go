package main

import (
	"fmt"
	"sync"
)

func main() {
	problem()
	// syncWithMutex()
	// copySolution()
}

func problem() {
	for {
		s := make([]int, 0, 1)
		go func() {
			s1 := append(s, 1)
			fmt.Printf("s1: %v\n", s1)
		}()
		go func() {
			s2 := append(s, 2)
			fmt.Printf("s2: %v\n", s2)
		}()
		fmt.Scanln()
	}
}

func syncWithMutex() {
	for {
		s := make([]int, 0, 1)
		mu := sync.Mutex{}
		go func() {
			mu.Lock()
			defer mu.Unlock()
			s1 := append(s, 1)
			fmt.Printf("s1: %v\n", s1)
		}()
		go func() {
			mu.Lock()
			defer mu.Unlock()
			s2 := append(s, 2)
			fmt.Printf("s2: %v\n", s2)
		}()
		fmt.Scanln()
	}
}

func copySolution() {
	for {
		s := make([]int, 0, 1)
		go func() {
			sCopy := make([]int, len(s), cap(s))
			copy(sCopy, s)
			s1 := append(sCopy, 1)
			fmt.Printf("s1: %v\n", s1)
		}()
		go func() {
			sCopy := make([]int, len(s), cap(s))
			copy(sCopy, s)
			s2 := append(sCopy, 2)
			fmt.Printf("s2: %v\n", s2)
		}()
		fmt.Scanln()
	}
}
