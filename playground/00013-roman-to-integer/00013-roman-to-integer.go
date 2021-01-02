package _0013_roman_to_integer

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	// min 0 max MMMM
	if s == "" {
		return 0
	}

	num := 0
	lastnum := 0
	total := 0
	// range string
	for i := 0; i < len(s); i++ {
		// first mid last roman
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastnum { // left roman
			total = total - num
		} else { // right roman
			total = total + num
		}
		lastnum = num // mark last
	}

	return total
}
