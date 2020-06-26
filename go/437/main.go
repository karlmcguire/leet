package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func fn(root *TreeNode, c, t int, s map[int]int) int {
	if root == nil {
		return 0
	}
	c += root.Val
	r, ok := s[c-t]
	if !ok {
		r = 0
	}
	a, ok := s[c]
	if !ok {
		a = 0
	}
	s[c] = a + 1
	r += fn(root.Left, c, t, s) + fn(root.Right, c, t, s)
	s[c] = s[c] - 1
	return r
}

func pathSum(root *TreeNode, sum int) int {
	s := make(map[int]int, 0)
	s[0] = 1
	return fn(root, 0, sum, s)
}

func main() {
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2}},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: -3,
			Right: &TreeNode{Val: 11}}}

	fmt.Printf("%v:3\n", pathSum(root, 8))
}
