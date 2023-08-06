package main

import "fmt"

func assertEqual(a, b []int) {
	err := fmt.Sprintf("a(%+v) != b(%+v)\n", a, b)
	if len(a) != len(b) {
		panic(err)
	}
	for i, v := range a {
		if b[i] != v {
			panic(err)
		}
	}
}

func main() {
	{
		s1 := []int{1, 2, 3}
		s2 := s1[1:2]
		s3 := append(s2, 10)

		assertEqual(s1, []int{1, 2, 10})
		assertEqual(s2, []int{2})
		assertEqual(s3, []int{2, 10})
	}
	{
		s1 := []int{1, 2, 3}
		s2 := s1[1:2:2]
		s3 := append(s2, 10)

		assertEqual(s1, []int{1, 2, 3})
		assertEqual(s2, []int{2})
		assertEqual(s3, []int{2, 10})
	}
}
