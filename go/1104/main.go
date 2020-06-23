package main

import (
	"fmt"
	"math"
)

// O(logn) runtime
func fn(label int) []int {
	// levels
	l := int(math.Floor(math.Log2(float64(label)))) + 1
	// output
	o := make([]int, l, l)
	// first element must be 1
	o[0] = 1
	// last element must be label
	o[l-1] = label

	// iterate through remaining slots
	for i := l - 2; i > 0; i-- {
		h := 1 << int(math.Log(float64(o[i+1]))/math.Log(2))
		o[i] = h - (o[i+1]^h)/2 - 1
	}

	return o
}

func main() {
	fmt.Printf("%v:[1, 3, 4, 14]\n", fn(14))
	fmt.Printf("%v:[1, 2, 6, 10, 26]\n", fn(26))
}
