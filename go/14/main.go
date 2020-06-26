package main

import (
	"bytes"
	"fmt"
	"sort"
)

func longestCommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	sort.Strings(s)
	a, b, o := s[0], s[len(s)-1], &bytes.Buffer{}
	for i := range a {
		if a[i] == b[i] {
			o.WriteByte(a[i])
		} else {
			break
		}
	}
	return o.String()
}

func main() {
	fmt.Printf("%v:fl\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%v:\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
