### 解题思路

- 按 `/` 分割字符串
- 遍历分割好的字符串，并注意边界条件
  - 遇到 `..` 要退回去
  - 遇到 `.` 并且有额外路径的，需要加上路径
- 最终如果 pathList 长度为 0 则返回根目录 `/`
- 其他情况拼接好 pathList

### 代码

```go
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	pathList := make([]string, 0)
	var res string
	for i := 0; i < len(arr); i++ {
		cur := strings.TrimSpace(arr[i])
		if cur == ".." {
			if len(pathList) > 0 {
				pathList = pathList[:len(pathList)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			pathList = append(pathList, cur)
		}
	}
	if len(pathList) == 0 {
		return "/"
	}
	res = strings.Join(pathList, "/")
	return "/" + res
}
```
