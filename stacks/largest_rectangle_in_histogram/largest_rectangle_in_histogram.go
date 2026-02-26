// Package largestrectangleinhistogram solves the largest rectangle in histogram problem.
package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := [][]int{} // [(index,height), (index, height)]
	for idx, h := range heights {
		start := idx
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			popIndex := len(stack) - 1
			value := stack[popIndex]
			stack = stack[:popIndex]
			if maxArea < (value[1] * (idx - value[0])) {
				maxArea = value[1] * (idx - value[0])
			}
			start = value[0]
		}
		stack = append(stack, []int{start, h})
	}
	for _, h := range stack {
		if maxArea < (h[1] * (len(heights) - h[0])) {
			maxArea = h[1] * (len(heights) - h[0])
		}
	}
	return maxArea
}
