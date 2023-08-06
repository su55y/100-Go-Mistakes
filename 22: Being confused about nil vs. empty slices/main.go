package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s1 []int
	s2 := []int(nil)
	s3 := []int{}
	s4 := make([]int, 0)

	assertTrue(
		len(s1) == 0 && s1 == nil,
		len(s2) == 0 && s2 == nil,
		len(s3) == 0 && s3 != nil,
		len(s4) == 0 && s4 != nil,
	)

	// nil and empty diff in encoding/json
	var operations1 []float32
	customer1 := customer{ID: "foo", Oparations: operations1}
	b, _ := json.Marshal(customer1)
	assertTrue(string(b) == `{"id":"foo","operations":null}`)

	operations2 := make([]float32, 0)
	customer2 := customer{ID: "bar", Oparations: operations2}
	b2, _ := json.Marshal(customer2)
	assertTrue(string(b2) == `{"id":"bar","operations":[]}`)

}

type customer struct {
	ID         string    `json:"id"`
	Oparations []float32 `json:"operations"`
}

func assertTrue(statements ...bool) {
	for i, s := range statements {
		if !s {
			panic(fmt.Sprintf("%d statement is false", i))
		}
	}
}
