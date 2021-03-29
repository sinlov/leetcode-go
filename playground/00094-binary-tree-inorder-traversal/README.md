### 解题思路

- 递归实现一个二叉树中序遍历

### 代码

```go
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
```
