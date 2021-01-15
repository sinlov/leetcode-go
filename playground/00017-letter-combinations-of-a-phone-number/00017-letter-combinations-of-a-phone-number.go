package _00017_letter_combinations_of_a_phone_number

var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	letterBiz("", digits)
	return res
}

func letterBiz(result string, digits string) {
	if digits == "" { // exit
		res = append(res, result)
		return
	}
	k := digits[0:1] // recall biz
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		result += dict[k][i] // back
		letterBiz(result, digits)
		result = result[0 : len(result)-1]
	}
}
