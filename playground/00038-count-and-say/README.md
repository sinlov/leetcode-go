### 解题思路

查看规律

| 值        | 表示                        |
|:----------|:---------------------------|
| 1         | 1 |
| 2         | 11 |
| 3         | 21 |
| 4         | 1211 |
| 5         | 111221 |
| 6         | 312211 |
| 7         | 13112221 |
| 8         | 1113213211 |

循环上一个数，并且标记第一个数字
当后面遇到第一个不为这个数字的时候, 把数字重复的次数转换成字符串，加到这个数字转换成的字符串上

双指针实现

- pre 记录前一项, 初始化为空， cur 为当前，初始化为 1
- 特殊情况，直接输出
- 定义指针 head, tail 先都指向序列的 头部
- tmp 为 游标缓存 初始化为 cur[0] ，length 初始化为 len(cur)
- 从 下标 1 遍历 到 n - 1
  - 每次遍历先重新初始化 length 初始化为 len(cur), tmp 缓存为 cur[0] head 指针重新付给开头
  - 尾指针，遍历 0 -> length
    - 如果 cur[tail] == tmp ，则 头指针自增
    - 如果 cur[tail] != tmp , 则 缓存到前一个记录，并将 tmp 移到 cur[tail] 位置，再将 head 放到 1 的位置，标记第一个值
    - 处理特殊情况 变量到最后一个 直接缓存到前一个记录
  - 双指针循环完毕后，将记录值 pre 赋值给 cur, 清空 pre 缓存，进入下一次循环
- 最终输出当前 cur

### 代码

```go
// 1 <= n <= 30
func countAndSay(n int) string {
	if n == 1 { // special case
		return "1"
	}
	pre, cur := "", "1"
	head, tail := 0, 0

	tmp := cur[0]
	length := len(cur)

	for i := 1; i < n; i++ { // 1 -> n -1
		length = len(cur) // reset length
		tmp = cur[0]      // reset tmp
		head = 0          // reset head and tail 0 0
		for tail = 0; tail < length; tail++ {
			if cur[tail] == tmp { // find same as cache
				head++
			} else { // not found add tmp and cache pre then mark head -> 1
				pre = pre + strconv.Itoa(head) + string(tmp)
				tmp = cur[tail]
				head = 1
			}

			if tail == length-1 { // special case length-1 cache pre
				pre = pre + strconv.Itoa(head) + string(tmp)
			}
		}
		cur = pre // cur now
		pre = ""  //reset
	}
	return cur
}
```
