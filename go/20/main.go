package main

import "fmt"

func isValid(s string) bool {
	b := make([]byte, 0)
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			b = append(b, s[i])
		} else {
			if len(b) == 0 {
				return false
			}
			if s[i] == ')' && b[len(b)-1] != '(' {
				return false
			}
			if s[i] == '}' && b[len(b)-1] != '{' {
				return false
			}
			if s[i] == ']' && b[len(b)-1] != '[' {
				return false
			}
			b = b[:len(b)-1]
		}
	}
	return len(b) == 0
}

func main() {
	fmt.Printf("%v:true\n", isValid("()"))
	fmt.Printf("%v:true\n", isValid("()[]{}"))
	fmt.Printf("%v:false\n", isValid("(]"))
	fmt.Printf("%v:false\n", isValid("([)]"))
	fmt.Printf("%v:true\n", isValid("{()}"))
}
