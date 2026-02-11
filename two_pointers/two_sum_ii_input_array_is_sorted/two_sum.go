// Package twosum will solve the Two Integer Sum II problem.
package twosum

func twoSum(nums []int, target int) []int {
	leftPointer, rightPointer := 0, len(nums)-1
	for leftPointer < rightPointer {
		leftValue := nums[leftPointer]
		rightValue := nums[rightPointer]
		if leftValue+rightValue == target {
			return []int{leftPointer + 1, rightPointer + 1}
		}
		if leftValue+rightValue > target {
			rightPointer--
			continue
		}
		leftPointer++
	}
	return []int{}
}
