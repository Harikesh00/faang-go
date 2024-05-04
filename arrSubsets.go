package main

// 41
func subsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int, subset []int)

	backtrack = func(start int, subset []int) {
		// Add the current subset to the result
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		result = append(result, subsetCopy)

		// Continue adding elements to the subset
		for i := start; i < len(nums); i++ {
			// Include the current element in the subset
			subset = append(subset, nums[i])

			// Recursively explore subsets with the current element included
			backtrack(i+1, subset)

			// Exclude the current element from the subset (backtrack)
			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0, []int{})
	return result
}

// func main() {
//     nums := []int{1, 2, 3}
//     fmt.Println(subsets(nums)) // Output: [[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]
// }
