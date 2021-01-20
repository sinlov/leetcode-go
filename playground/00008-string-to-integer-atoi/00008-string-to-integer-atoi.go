package _00008_string_to_integer_atoi

import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	strpStr := stripLeftWhitespace(s)
	positive := true
	res := 0
	for i, char := range strpStr {
		num := int(char - '0')
		if i == 0 { // check symbol
			if char == '-' {
				positive = false
			} else if char == '+' {
				positive = true
			} else if num >= 0 && num <= 9 {
				res = res*10 + num
			} else {
				return res
			}
		} else {
			if num < 0 || num > 9 { // not number
				break
			}
			if positive && res*10+num > math.MaxInt32 { // max
				return math.MaxInt32
			}
			if !positive && res*10-num < math.MinInt32 {// min
				return math.MinInt32
			}
			if positive {
				res = res*10 + num
			} else {
				res = res*10 - num
			}
		}
	}
	return res
}

func stripLeftWhitespace(str string) string {
	res := ""
	headNotWhitespace := true
	for _, char := range str {
		if headNotWhitespace {
			if char != ' ' {
				headNotWhitespace = false
				res = res + string(char)
			}
		} else {
			res = res + string(char)
		}
	}
	return res
}
