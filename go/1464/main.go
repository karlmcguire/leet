package main

import "fmt"

// O(n) runtime
// O(1) storage
func maxProduct(n []int) int {
	m := []int{0, 0}
	for i := range n {
		if n[i] > m[0] {
			m[0], m[1] = n[i], m[0]
		} else if n[i] > m[1] {
			m[1] = n[i]
		}
	}
	return (m[0] - 1) * (m[1] - 1)
}

func main() {
	fmt.Printf("%v:12\n", maxProduct([]int{3, 4, 5, 2}))
	fmt.Printf("%v:16\n", maxProduct([]int{1, 5, 4, 5}))
	fmt.Printf("%v:12\n", maxProduct([]int{3, 7}))
}
