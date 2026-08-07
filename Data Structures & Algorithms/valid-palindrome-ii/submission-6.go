func validPalindrome(s string) bool {
		mid := len(s) / 2
	i := 0
	j := len(s) - 1
	for i < mid {
		if s[i] != s[j] {
			// Try deleting a char
			left := checkIfSubstringIsPalindrome(s[i+1 : j+1])
			right := checkIfSubstringIsPalindrome(s[i:j])
			if !left && !right {
				return false
			}
		}
		i++
		j--
	}
	return true
}

func checkIfSubstringIsPalindrome(s string) bool {
	mid := len(s) / 2
	i := 0
	j := len(s) - 1

	for i < mid {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}