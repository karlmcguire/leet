package main

import "fmt"

func strStr(h string, n string) int {
	if n == "" || n == h {
		return 0
	}
	if len(n) > len(h) {
		return -1
	}
global:
	for i := 0; i < len(h); i++ {
		if h[i] != n[0] {
			continue
		}
		if len(n) == 1 {
			return i
		}
		for j := 1; j < len(n) && i+j < len(h); j++ {
			if h[i+j] != n[j] {
				continue global
			}
			if j == len(n)-1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Printf("%v:2\n", strStr("hello", "ll"))
	fmt.Printf("%v:-1\n", strStr("aaaaa", "bba"))
	fmt.Printf("%v:0\n", strStr("aaaaa", "aa"))
	fmt.Printf("%v:-1\n", strStr("mississippi", "issipi"))
	fmt.Printf("%v:0\n", strStr("aaa", "a"))
}
