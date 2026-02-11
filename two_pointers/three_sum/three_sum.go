// Package threesum will solve the threesum problem.
package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		firstValue := nums[i]
		// The lowest value greater than zero, there is no way to make zero.
		// The array is sorted.
		if firstValue > 0 {
			break
		}
		// Check for a duplicate firstValue.
		if i > 0 && firstValue == nums[i-1] {
			continue
		}
		// firstValue[i], while leftPointer will be i+1 and rightPointer at the end.
		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := firstValue + nums[l] + nums[r]
			// Value is too big, move right pointer.
			if threeSum > 0 {
				r--
				// Value is too low, move left pointer.
			} else if threeSum < 0 {
				l++
			} else {
				// It's Zero so append the three values.
				res = append(res, []int{firstValue, nums[l], nums[r]})
				l++ // move pointer
				r-- // move pointer.
				// move left pointer to prevent duplicate values
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
