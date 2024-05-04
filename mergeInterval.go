package main

import (
	"sort"
)

// 34
func merge(intervals [][]int) [][]int {
	// Sort intervals based on the start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		// Check if the current interval overlaps with the previous merged interval
		if intervals[i][0] <= merged[len(merged)-1][1] {
			// Update the end time of the previous merged interval if needed
			if intervals[i][1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			}
		} else {
			// If there is no overlap, add the current interval to the merged list
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

// func main() {
//     intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
//     fmt.Println(merge(intervals)) // Output: [[1 6] [8 10] [15 18]]
// }
