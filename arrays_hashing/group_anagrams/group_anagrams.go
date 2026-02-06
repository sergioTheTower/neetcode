// Package groupanagrams solves the Group Anagrams problem.
package groupanagrams

func GroupAnagrams(strs []string) [][]string {
	resultMap := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int // Fixed-size array initialized to zeros
		for i := 0; i < len(s); i++ {
			// In Go, byte subtraction gives the index
			count[s[i]-'a']++
		}
		// Use the array as the key directly
		resultMap[count] = append(resultMap[count], s)
	}
	// Pull out the values from the map into a slice.
	output := make([][]string, 0, len(resultMap))
	for _, group := range resultMap {
		output = append(output, group)
	}
	return output
}
