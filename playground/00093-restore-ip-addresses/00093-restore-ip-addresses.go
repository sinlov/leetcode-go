package _00093_restore_ip_addresses

import "strconv"

func restoreIpAddresses(s string) []string {
	if s == "" { // special case
		return []string{}
	}
	res, ip := []string{}, []int{}
	dfsIPAddress(s, 0, ip, &res)
	return res
}

func dfsIPAddress(s string, index int, ip []int, res *[]string) {
	if index == len(s) {
		if len(ip) == 4 {
			*res = append(*res, intIPList2String(ip))
		}
		return
	}
	if index == 0 {
		num, _ := strconv.Atoi(string(s[0]))
		ip = append(ip, num)
		dfsIPAddress(s, index+1, ip, res)
	} else {
		num, _ := strconv.Atoi(string(s[index]))
		next := ip[len(ip)-1]*10 + num
		if next <= 255 && ip[len(ip)-1] != 0 {
			ip[len(ip)-1] = next
			dfsIPAddress(s, index+1, ip, res)
			ip[len(ip)-1] /= 10
		}
		if len(ip) < 4 {
			ip = append(ip, num)
			dfsIPAddress(s, index+1, ip, res)
			ip = ip[:len(ip)-1]
		}
	}
}

func intIPList2String(ip []int) string {
	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}
