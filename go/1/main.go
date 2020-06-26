package main

import "fmt"

// O(n^2) - naive
/*
func twoSum(n []int, t int) []int {
	var s [2]int
	for i := range n {
		for j := range n {
			if i != j && n[i]+n[j] == t {
				s[0], s[1] = i, j
			}
		}
	}
	return s[:]
}
*/

// O(n)
func twoSum(n []int, t int) []int {
	b := make(map[int]int, len(n))
	for i := range n {
		if j, ok := b[t-n[i]]; ok {
			return []int{j, i}
		}
		b[n[i]] = i
	}
	return nil
}

func main() {
	fmt.Printf("%v:[0 1]\n", twoSum([]int{2, 7, 11, 15}, 9))
}
