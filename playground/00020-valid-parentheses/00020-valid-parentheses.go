package _00020_valid_parentheses

func isValid(s string) bool {
	// empty is true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') { // push
			stack = append(stack, v)
		} else {

			if (v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') || // valid ()
				(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') || // valid [ ]
				(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') { // valid { }
				stack = stack[:len(stack)-1]
			} else {
				return false // not valid
			}
		}
	}

	return len(stack) == 0 // final 0 is valid
}
