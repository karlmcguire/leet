package main

import "fmt"

type bitset []uint64

func newBitset(size uint64) bitset { return make(bitset, (size+63)/64) }
func (s bitset) set(b int)         { s[b/64] |= 1 << (b & 63) }
func (s bitset) get(b int) bool    { return s[b/64]&(1<<(b&63)) != 0 }

// O(n) runtime
// O(1) memory
func fn(A, B []int) []int {
	s := newBitset(200002)
	a := 0
	b := 0

	for i := range A {
		a += A[i]
	}
	for i := range B {
		b += B[i]
		s.set(2 * B[i])
	}

	d := a - b

	for i := range A {
		e := 2*A[i] - d
		if e > 0 && e < 200002 && s.get(e) {
			return []int{A[i], (2*A[i] - d) / 2}
		}
	}

	return []int{}
}

func main() {
	fmt.Printf("%v:[1, 2]\n", fn([]int{1, 1}, []int{2, 2}))
	fmt.Printf("%v:[1, 2]\n", fn([]int{1, 2}, []int{2, 3}))
	fmt.Printf("%v:[2, 3]\n", fn([]int{2}, []int{1, 3}))
	fmt.Printf("%v:[5, 4]\n", fn([]int{1, 2, 5}, []int{2, 4}))
}
