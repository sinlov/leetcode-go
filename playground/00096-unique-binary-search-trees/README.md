### 解题思路

遇到次数第一反应 `动态规划`

- dp[n] 代表 1-n 个数能组成多少个不同的二叉排序树
- F(i,n) 代表以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数

那么
```
dp[n] = F(1,n) + F(2,n) + F(3,n) + …… + F(n,n)
```

- 初始值 dp[0] = 1，dp[1] = 1
- 分析 dp 和 F(i,n) 的关系又可以得到下面这个等式
```
F(i,n) = dp[i-1] * dp[n-i]
```

> 由于二叉排序树本身的性质，右边的子树一定比左边的子树，值都要大。所以这里只需要根节点把树分成左右，不需要再关心左右两边数字的大小

- 状态转移方程
```
dp[i] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + …… + dp[n-1] * dp[0]
```

- 最终得到的结果为 dp[n]

### 代码

```go
func numTrees() {
}
```
