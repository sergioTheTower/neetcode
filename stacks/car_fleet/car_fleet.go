package carfleet

import "sort"

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	pair := make([][2]int, n)
	for i := 0; i < n; i++ {
		pair[i] = [2]int{position[i], speed[i]}
	}
	// Reverse Order of the slice.
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0]
	})

	fleets := 1
	prevTime := float64(target-pair[0][0]) / float64(pair[0][1])
	for i := 1; i < n; i++ {
		currTime := float64(target-pair[i][0]) / float64(pair[i][1])
		if currTime > prevTime {
			fleets++
			prevTime = currTime
		}
	}
	return fleets
}
