package main

import "fmt"

func unpredictableMutation(m map[int]bool) map[int]bool {
	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	return m
}

func predictableMutation(m map[int]bool) map[int]bool {
	m2 := copyMap(m)
	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}
	return m2
}

func copyMap(m map[int]bool) map[int]bool {
	mcopy := make(map[int]bool, len(m))
	for k, v := range m {
		mcopy[k] = v
	}
	return mcopy
}

func main() {
	unpredictable := unpredictableMutation(
		map[int]bool{
			0: true,
			1: false,
			2: true,
		})

	predictable := predictableMutation(
		map[int]bool{
			0: true,
			1: false,
			2: true,
		})

	fmt.Println(unpredictable)
	fmt.Println(predictable)
}
