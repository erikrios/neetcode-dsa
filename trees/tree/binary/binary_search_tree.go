package binary

import "trees/tree"

type BinarySearchTree struct {
	Root *tree.TreeNode
}

func New() *BinarySearchTree {
	return &BinarySearchTree{Root: nil}
}

func (b *BinarySearchTree) Search(val int) bool {
	return search(b.Root, val)
}

func (b *BinarySearchTree) Insert(val int) {
	if b.Root == nil {
		node := tree.New()
		node.Val = val
		b.Root = node
	} else {
		insert(b.Root, val)
	}
}

func (b *BinarySearchTree) Delete(val int) {
	delete(b.Root, val)
}

func minValueNode(root *tree.TreeNode) *tree.TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
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

func insert(root *tree.TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			node := tree.New()
			node.Val = val
			root.Left = node
			return
		}
		insert(root.Left, val)
	} else {
		if root.Right == nil {
			node := tree.New()
			node.Val = val
			root.Right = node
			return
		}
		insert(root.Right, val)
	}
}

func delete(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		root.Right = delete(root.Right, val)
	} else if val < root.Val {
		root.Left = delete(root.Left, val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := minValueNode(root.Right)
			root.Val = minNode.Val
			root.Right = delete(root.Right, minNode.Val)
		}
	}

	return root
}
