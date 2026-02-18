// Package trapwater will solve the trap_rain_water question.
package trapwater

func trapWater(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	result := 0
	l, r := 0, len(heights)-1
	lMax, rMax := heights[l], heights[r]
	for l < r {
		if lMax < rMax {
			l += 1
			lMax = max(lMax, heights[l])
			result += lMax - heights[l]
			continue
		}
		r -= 1
		rMax = max(rMax, heights[r])
		result += rMax - heights[r]

	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
