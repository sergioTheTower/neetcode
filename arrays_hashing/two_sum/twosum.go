package twosum

func TwoSum(nums []int, target int) []int {
	// The key is the diff, the value is the index.
	trackDiffs := map[int]int{}
	for i, v := range nums {
		diff := target - v
		if idx, ok := trackDiffs[diff]; ok {
			return []int{idx, i}
		}
		trackDiffs[v] = i
	}
	return []int{}
}
