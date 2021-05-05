package _00101_symmetric_tree

import "github.com/sinlov/leetcode-go/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func isSymmetric(root *TreeNode) bool {
	return root == nil || dfs(root.Left, root.Right)
}

/**
* DFS compare tree node
 */
func dfs(rootLeft, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil { //  both nil
		return true
	}
	if rootLeft == nil || rootRight == nil {
		return false
	}
	if rootLeft.Val != rootRight.Val {
		return false
	}
	return dfs(rootLeft.Left, rootRight.Right) && dfs(rootLeft.Right, rootRight.Left) // child node
}
