// Package topkfrequent will solve the Top K Frequent Elements problem.
package topkfrequent

func topKFrequent(nums []int, k int) []int {
	// Get the int freq count with a hash map.
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	// Use bucket sort but use the index of the first array to represent the count and
	// The int value is appended to the nested array.
	freqBuckets := make([][]int, len(nums)+1)
	for num, cnt := range count {
		freqBuckets[cnt] = append(freqBuckets[cnt], num)
	}
	res := []int{}
	for i := range freqBuckets {
		// Loop from the right to left using len(freqBuckets)-i-1.
		for _, num := range freqBuckets[len(freqBuckets)-i-1] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
