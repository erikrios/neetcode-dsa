package binary

import "trees/tree"

type BinarySearchTree struct {
	Root *tree.TreeNode
}

func New(val int) *BinarySearchTree {
	return &BinarySearchTree{Root: tree.New(val)}
}

func Search(root *tree.TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if val < root.Val {
		return Search(root.Left, val)
	} else if val > root.Val {
		return Search(root.Right, val)
	} else {
		return true
	}
}
