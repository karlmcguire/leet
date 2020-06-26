package main

import "fmt"

func maxSatisfied(customers, grumpy []int, X int) int {
	s, m, t := 0, 0, 0
	for i := range customers {
		if grumpy[i] == 1 {
			t += customers[i]
		} else {
			s += customers[i]
		}
		if i >= X && grumpy[i-X] == 1 {
			t -= customers[i-X]
		}
		if t > m {
			m = t
		}
	}
	return s + m
}

func main() {
	fmt.Printf("%v:16\n", maxSatisfied(
		[]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1},
		3,
	))
	fmt.Printf("%v:1\n", maxSatisfied(
		[]int{1},
		[]int{0},
		1,
	))
	fmt.Printf("%v:24\n", maxSatisfied(
		[]int{4, 10, 10},
		[]int{1, 1, 0},
		2,
	))
	fmt.Printf("%v:29\n", maxSatisfied(
		[]int{7, 8, 8, 6},
		[]int{0, 1, 0, 1},
		3,
	))
}
