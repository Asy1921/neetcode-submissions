func longestCommonPrefix(strs []string) string {
    longestPrefix := ""
	for i := 0; i < len(strs[0]); i++ {
		currentChar := (strs[0][i])
		// Check if the longestPrefix is still valid
		stillValid := true
		for j := 1; j < len(strs); j++ {
			// Loop over the strs and check if prefix is still valid

			if !(len(strs[j]) > i && currentChar == strs[j][i]) {
				stillValid = false
			}

		}
		if stillValid {
			longestPrefix = longestPrefix + string(currentChar)
		} else {
			return longestPrefix
		}

	}
	return longestPrefix
}
