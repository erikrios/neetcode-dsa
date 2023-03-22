package main

import (
	"fmt"
)

func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
	}

	paths := make([]int, 0)
	fmt.Println(canReachLeaf(node, &paths))
	fmt.Println(paths)

	node = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 0,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	paths = make([]int, 0)
	fmt.Println(canReachLeaf(node, &paths))
	fmt.Println(paths)
}

func canReachLeaf(root *TreeNode, paths *[]int) bool {
	if root == nil || root.Val == 0 {
		return false
	}

	*paths = append(*paths, root.Val)

	if root.Left == nil && root.Right == nil {
		return true
	}

	lVal := canReachLeaf(root.Left, paths)
	rVal := canReachLeaf(root.Right, paths)
	res := lVal || rVal
	if !res {
		temps := make([]int, len(*paths)-1, len(*paths)-1)

		for i := 0; i < len(temps); i++ {
			temps[i] = (*paths)[i]
		}

		*paths = temps
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New() *TreeNode {
	return &TreeNode{}
}
