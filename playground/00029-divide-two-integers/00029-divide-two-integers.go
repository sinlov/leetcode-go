package _00029_divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	res := 0 // quotient start at 0

	symbol := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 { // final symbol
		symbol = 1
	}

	absDivd, absDivs := abs(dividend), abs(divisor)

	for absDivd >= absDivs { // dichotomization
		tmp := absDivs
		mid := 1
		for tmp<<1 <= absDivd { // out (tmp << 1) <= dividend
			tmp <<= 1
			mid <<= 1
		}
		absDivd -= tmp
		res += mid
	}

	return symbol * res // final add symbol
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
