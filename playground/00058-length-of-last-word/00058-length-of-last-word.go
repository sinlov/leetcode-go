package _00058_length_of_last_word

func lengthOfLastWord(s string) int {
	if len(s) == 0 { // speace case
		return 0
	}
	tail := len(s) - 1
	for tail >= 0 && s[tail] == ' ' { // find out tail not space
		tail--
	}
	head := tail
	for head >= 0 && s[head] != ' ' { // find out head space
		head--
	}
	return tail - head
}
