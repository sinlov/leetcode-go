### 解题思路

- 将根节点的左子树反转二叉树
- 再和根节点的右节点进行比较
  - 相等则为完全二叉树
  - 否则不是

> 反转二叉树是第 226 题

### 代码

```go
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
```
