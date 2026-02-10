// Package lcs contains functions for finding the longest consecutive sequence.
package lcs

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}
	longest := 0
	for num := range set {
		// Check if the next number exist.
		length := 1
		for set[num+length] {
			// Increase every time a new consecutive number is found.
			length++
		}
		// No more consecutive values were found, check if there is a new longest
		// consecutive numbers.
		if length > longest {
			longest = length
		}
	}
	return longest
}
