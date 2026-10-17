func gcdOfStrings(str1 string, str2 string) string {
	maxLen := len(str1)
	if len(str2) < maxLen {
		maxLen = len(str2)
	}

	for candidateLen := maxLen; candidateLen > 0; candidateLen-- {
		if len(str1)%candidateLen != 0 || len(str2)%candidateLen != 0 {
			continue
		}

		candidate := str1[:candidateLen]
		if strings.Repeat(candidate, len(str1)/candidateLen) == str1 &&
			strings.Repeat(candidate, len(str2)/candidateLen) == str2 {
			return candidate
		}
	}

	return ""
}