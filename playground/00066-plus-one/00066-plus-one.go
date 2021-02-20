package _00066_plus_one

func plusOne(digits []int) []int {
	length := len(digits)
	carry := 1 // default carry 1
	for i := length - 1; i >= 0; i-- {
		if carry == 0 { // no need carry break
			break
		}
		if digits[i]+carry > 9 { // need carry on
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	if digits[0] == 0 && carry == 1 { // subscript 0 need carry on
		digits = append([]int{1}, digits...)
	}
	return digits
}
