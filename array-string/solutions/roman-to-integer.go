package solutions

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {
	var result int = 0
	for i, val := range s {
		v := m[string(val)]
		if len(s) > i+1 {
			v1 := m[string(s[i+1])]
			if v >= v1 {
				result += m[string(s[i])]
			} else {
				result -= m[string(s[i])]
			}
		} else {
			result += m[string(s[i])]
		}
	}
	return result
}
