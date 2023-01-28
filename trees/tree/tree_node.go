package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New() *TreeNode {
	return &TreeNode{}
}
