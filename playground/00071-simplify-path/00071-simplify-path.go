package _00071_simplify_path

import "strings"

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
