package _00043_multiply_strings

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" { // special case
		return "0"
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1 // more bigger numer to front
	}

	cacheMap := make(map[int]int)

	// let converts to binary operation
	for i, v := range num1 {
		v = v - 48 // ASCII
		vIndex := len(num1) - i - 1
		for j, k := range num2 {
			k = k - 48 // ASCII
			kIndex := len(num2) - j - 1

			resIndex := vIndex + kIndex
			resV := v * k

			cacheMap[resIndex] = int(resV) + cacheMap[resIndex]
		}
	}

	// max length 200
	for i := 0; i < 200; i++ {
		if v := cacheMap[i]; v > 9 {
			carry := v / 10
			cacheMap[i] = v % 10
			cacheMap[i+1] = cacheMap[i+1] + carry
		}
	}

	// converts to decimalism
	res := ""
	isPassLeftZero := false
	// max length 200
	for i := 199; i >= 0; i-- {
		v := cacheMap[i]
		if !isPassLeftZero && v == 0 {
			continue
		} else {
			if !isPassLeftZero && v != 0 {
				isPassLeftZero = true
			}
			res = res + strconv.Itoa(v)
		}
	}
	return res
}
