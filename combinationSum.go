package main

// 32
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func(combination []int, startIndex, target int)

	backtrack = func(combination []int, startIndex, target int) {
		if target == 0 {
			// If target is achieved, add the combination to the result
			result = append(result, append([]int{}, combination...))
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			// If the current candidate is larger than the target, skip it
			if candidates[i] > target {
				continue
			}

			// Add the current candidate to the combination
			combination = append(combination, candidates[i])

			// Recursively explore combinations with the current candidate
			backtrack(combination, i, target-candidates[i])

			// Backtrack by removing the last added candidate
			combination = combination[:len(combination)-1]
		}
	}

	backtrack([]int{}, 0, target)
	return result
}

// func main() {
//     candidates := []int{2, 3, 6, 7}
//     target := 7
//     fmt.Println(combinationSum(candidates, target)) // Output: [[2 2 3] [7]]
// }
