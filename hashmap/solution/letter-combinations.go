package solution

import "fmt"

var phoneDigits = map[string][]string{
	"1": {},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
	"":  {},
}

func LetterCombinations(digits string) []string {
	if digits == "" || len(digits) == 1 {
		return phoneDigits[digits]
	}

	var combination []string

	for _, val1 := range phoneDigits[string(digits[0])] {
		for i := 1; i < len(digits); i++ {
			for _, val2 := range phoneDigits[string(digits[i])] {
				combination = append(combination, fmt.Sprintf("%s%s", val1, val2))
			}

		}
	}
	return combination
}
