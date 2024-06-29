package solution

func CanConstruct(ransomNote string, magazine string) bool {
	var mr = map[rune]int{}
	for _, val := range magazine {
		if _, ok := mr[val]; ok {
			mr[val]++
			continue
		}
		mr[val] = 1
	}
	for _, val := range ransomNote {
		if sr, ok := mr[val]; ok {
			if sr > 0 {
				mr[val]--
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
