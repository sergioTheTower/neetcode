// Package productofarrayexceptself provides a function to calculate the product of an array excluding each element.
package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for idx := range result {
		result[idx] = 1
	}

	prefix := 1
	for i := range nums {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := range nums {
		end := len(nums) - 1 - i
		result[end] *= postfix
		postfix *= nums[end]
	}
	return result
}
