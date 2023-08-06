package main

import (
	"crypto/rand"
	"fmt"
	"runtime"
)

func randomBytes() *[128]byte {
	var bytes [128]byte
	_, _ = rand.Read(bytes[:])
	return &bytes
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
}

func main() {
	n := 1_000_000

	// 461 MB after fill, 293 MB after delete
	// m := make(map[int][128]byte)

	// 182 MB after fill, 38 MB after delete
	m := make(map[int]*[128]byte)

	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = randomBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}
