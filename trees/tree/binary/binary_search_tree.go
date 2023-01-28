package binary

import "trees/tree"

type BinarySearchTree struct {
	Root *tree.TreeNode
}

func New(val int) *BinarySearchTree {
	return &BinarySearchTree{Root: tree.New()}
}

func (b *BinarySearchTree) Search(val int) bool {
	return search(b.Root, val)
}

func search(root *tree.TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if val < root.Val {
		return search(root.Left, val)
	} else if val > root.Val {
		return search(root.Right, val)
	} else {
		return true
	}
}
