package _00067_add_binary

import "strconv"

func addBinary(a string, b string) string {
	// let len(a) > len(b)
	if len(a) < len(b) {
		a, b = b, a
	}
	var complementZeroLength = len(a) - len(b)
	if len(a) != len(b) { // complement zero
		for i := 0; i < complementZeroLength; i++ {
			b = "0" + b
		}
	}

	var carry = 0 // carry flat if 1 need carry on
	var res = ""
	for i := len(a) - 1; i >= 0; i-- { // reverse
		aInt := a[i] - '0'
		bInt := b[i] - '0'
		sum := int(aInt+bInt) + carry
		if sum == 3 { // need carry on
			res = "1" + res
			carry = 1
		} else if sum == 2 { // need carry on
			res = "0" + res
			carry = 1
		} else {
			res = strconv.Itoa(sum) + res
			carry = 0
		}
	}

	if carry == 1 { // last flat need to carry on
		res = "1" + res
	}

	return res
}
