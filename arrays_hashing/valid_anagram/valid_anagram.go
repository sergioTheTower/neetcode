package validanagram

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := map[byte]int{}
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
		charCount[t[i]]--
	}
	for _, value := range charCount {
		if value != 0 {
			return false
		}
	}
	return true
}
