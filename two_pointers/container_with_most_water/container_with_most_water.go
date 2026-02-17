// Package containswater will find the area of the greatest amount of water stored.
package containswater

func containsWater(heights []int) int {
	var result int
	left, right := 0, len(heights)-1

	for left < right {
		leftValue := heights[left]
		rightValue := heights[right]
		// Compute area, right is the width and height is the left or right value.
		area := (right - 1) * min(leftValue, rightValue)

		if area > result {
			result = area
		}
		if leftValue < rightValue {
			left += 1
			continue
		}
		right -= 1
	}

	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
