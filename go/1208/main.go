package main

import "fmt"

func abs(n int) int { return (n ^ (n >> 31)) - (n >> 31) }

func equalSubstring(s, t string, k int) int {
	i, j := 0, 0
	for ; i < len(s); i++ {
		if k -= abs(int(s[i]) - int(t[i])); k < 0 {
			k, j = k+abs(int(s[j])-int(t[j])), j+1
		}
	}
	return i - j
}

func main() {
	fmt.Printf("%v:3\n", equalSubstring("abcd", "bcdf", 3))
	fmt.Printf("%v:1\n", equalSubstring("abcd", "cdef", 3))
	fmt.Printf("%v:1\n", equalSubstring("abcd", "acde", 0))
}
