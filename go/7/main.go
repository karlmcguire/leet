package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	s := []byte(fmt.Sprintf("%d", x))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == '-' {
			j++
			continue
		}
		s[i], s[j] = s[j], s[i]
	}
	n, err := strconv.ParseInt(string(s), 10, 32)
	if err != nil {
		return 0
	}
	return int(n)
}

func main() {
	fmt.Printf("%v:321\n", reverse(123))
	fmt.Printf("%v:-321\n", reverse(-123))
	fmt.Printf("%v:21\n", reverse(120))
	fmt.Printf("%v:0\n", reverse(1534236469))
}
