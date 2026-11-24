package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Zigzag returns level order with alternating direction, writing each level
// directly into a pre-sized slice: leftward levels fill from the end.
func Zigzag(root *TreeNode) [][]int {
	out := [][]int{}
	if root == nil {
		return out
	}
	queue := []*TreeNode{root}
	leftToRight := true
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			if leftToRight {
				level[i] = n.Val
			} else {
				level[size-1-i] = n.Val
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		out = append(out, level)
		leftToRight = !leftToRight
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	fmt.Println(Zigzag(root))
}
