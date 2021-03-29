package _00094_binary_tree_inorder_traversal

import "github.com/sinlov/leetcode-go/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, out *[]int) {
	if root != nil {
		inorder(root.Left, out)
		*out = append(*out, root.Val)
		inorder(root.Right, out)
	}
}
