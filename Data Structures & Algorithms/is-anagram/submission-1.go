func isAnagram(s string, t string) bool {
	if(len(s)!=len(t)){
		return false
	}
	mapS, mapT := make(map[rune]int), make(map[rune]int)
	// Map out count of each character
	for i:=0;i<len(s);i++{
		mapS[rune(s[i])]++;
		mapT[rune(t[i])]++;
	}

	// Compare freq of each char
	for i:=range(mapS){
		if mapS[i]!=mapT[i]{
			return false
		}
	}
	return true
}
