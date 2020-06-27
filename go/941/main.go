package main

import "fmt"

func validMountainArray(n []int) bool {
	if len(n) < 3 {
		return false
	}
	d := false
	for i, l := 1, n[0]; i < len(n); i, l = i+1, n[i] {
		if !d && l > n[i] {
			if i == 1 {
				return false
			}
			d = true
			continue
		}
		if l == n[i] {
			return false
		}
		if d && l < n[i] {
			return false
		}
	}
	return d
}

func main() {
	fmt.Printf("%v:false\n", validMountainArray([]int{2, 1}))
	fmt.Printf("%v:false\n", validMountainArray([]int{3, 5, 5}))
	fmt.Printf("%v:true\n", validMountainArray([]int{0, 3, 2, 1}))
	fmt.Printf("%v:false\n", validMountainArray([]int{0, 2, 3, 3, 5, 2, 1, 0}))
	fmt.Printf("%v:false\n", validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Printf("%v:false\n", validMountainArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}
