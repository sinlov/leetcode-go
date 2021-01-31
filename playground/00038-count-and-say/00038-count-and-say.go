package _00038_count_and_say

import "strconv"

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
