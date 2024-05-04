package main

import (
	"sort"
)

// 48
type Job struct {
	start, end, profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]Job, n)

	// Construct Job array
	for i := 0; i < n; i++ {
		jobs[i] = Job{startTime[i], endTime[i], profit[i]}
	}

	// Sort jobs based on end time
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	// Define dp array to store maximum profit
	dp := make([]int, n)
	dp[0] = jobs[0].profit

	for i := 1; i < n; i++ {
		// Find the latest job that finishes before the start time of the current job
		prev := binarySearch(jobs, i)

		// Calculate maximum profit considering the current job and the previous maximum profit
		includingProfit := jobs[i].profit
		if prev != -1 {
			includingProfit += dp[prev]
		}

		// Update dp array with the maximum profit
		dp[i] = max(dp[i-1], includingProfit)
	}

	return dp[n-1]
}

// Binary search to find the latest job that finishes before the start time of the current job
func binarySearch(jobs []Job, index int) int {
	left, right := 0, index-1
	for left <= right {
		mid := left + (right-left)/2
		if jobs[mid].end <= jobs[index].start {
			if jobs[mid+1].end <= jobs[index].start {
				left = mid + 1
			} else {
				return mid
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	startTime := []int{1, 2, 3, 3}
// 	endTime := []int{3, 4, 5, 6}
// 	profit := []int{50, 10, 40, 70}
// 	fmt.Println(jobScheduling(startTime, endTime, profit)) // Output: 120
// }
