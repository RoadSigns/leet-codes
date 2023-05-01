package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	var l *TreeNode
	var r *TreeNode

	if root == nil {
		return root
	}

	if root.Left != nil {
		l = InvertTree(root.Left)
	}

	if root.Right != nil {
		r = InvertTree(root.Right)
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  r,
		Right: l,
	}
}
