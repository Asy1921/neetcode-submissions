func mergeAlternately(word1 string, word2 string) string {
	res := ""
	i := 0
	j := 0
	for i < len(word1) || j < len(word2) {
		// Make sure we are not out of bounds
		if i < len(word1) {
			res += string(word1[i])
			i++
		}
		if j < len(word2) {
			res += string(word2[j])
			j++
		}

	}
	return res

}