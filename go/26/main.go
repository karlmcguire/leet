package main

import "fmt"

func removeDuplicates(n []int) int {
	if n == nil || len(n) == 0 {
		return 0
	}
	i, j, l := 1, 1, n[0]
	for ; i < len(n); i++ {
		if n[i] != l {
			n[j], j, l = n[i], j+1, n[i]
		}
	}
	return j
}

func main() {
	a := []int{1, 1, 2}
	fmt.Printf("%v:[1 2]\n", a[:removeDuplicates(a)])
	b := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("%v:[0 1 2 3 4]\n", b[:removeDuplicates(b)])
	c := []int{}
	fmt.Printf("%v\n", c[:removeDuplicates(c)])
	d := []int{1}
	fmt.Printf("%v\n", d[:removeDuplicates(d)])
}
