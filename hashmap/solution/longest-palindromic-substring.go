package solution

func LongestPalindrome(s string) string {
	longest := s[:1]
	for i := 0; i < len(s); i++ {
		temp := s[i:]
		for j := 0; j < len(s)-i; j++ {
			subStr := temp[:len(s)-j-i]
			if palindrome(subStr) && len(subStr) > len(longest) {
				longest = subStr
			}
		}
	}

	return longest
}

func palindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
