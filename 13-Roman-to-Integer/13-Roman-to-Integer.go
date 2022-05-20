package RomanToInteger

var romanLetters = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	currentNum, lastNum, result := 0, 0, 0
	s_len := len(s)
	for i := 0; i < s_len; i++ {
		character := s[s_len-(i+1) : s_len-i]
		currentNum = romanLetters[character]
		if currentNum < lastNum {
			result = result - currentNum
		} else {
			result = result + currentNum
		}
		lastNum = currentNum
	}
	return result
}
