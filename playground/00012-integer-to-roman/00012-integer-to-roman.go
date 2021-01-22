package _00012_integer_to_roman

func intToRoman(num int) string {
	res := ""
	copy := num
	for copy != 0 {
		if copy >= 1000 {
			copy = copy - 1000
			res += "M"
		} else if copy >= 900 {
			copy -= 900
			res += "CM"
		} else if copy >= 500 {
			copy -= 500
			res += "D"
		} else if copy >= 400 {
			copy -= 400
			res += "CD"
		} else if copy >= 100 {
			copy -= 100
			res += "C"
		} else if copy >= 90 {
			copy -= 90
			res += "XC"
		} else if copy >= 50 {
			copy -= 50
			res += "L"
		} else if copy >= 40 {
			copy -= 40
			res += "XL"
		} else if copy >= 10 {
			copy -= 10
			res += "X"
		} else if copy >= 9 {
			copy -= 9
			res += "IX"
		} else if copy >= 5 {
			copy -= 5
			res += "V"
		} else if copy >= 4 {
			copy -= 4
			res += "IV"
		} else if copy >= 1 {
			copy -= 1
			res += "I"
		}
	}
	return res
}
