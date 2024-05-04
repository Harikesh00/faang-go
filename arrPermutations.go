package main

// 33
func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(nums) {
			// If we've reached the end of the array, add the current permutation to the result
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		for i := start; i < len(nums); i++ {
			// Swap the current element with the element at index 'start'
			nums[start], nums[i] = nums[i], nums[start]

			// Recursively explore permutations starting from the next index
			backtrack(start + 1)

			// Undo the swap to backtrack
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	return result
}

// func main() {
//     nums := []int{1, 2, 3}
//     fmt.Println(permute(nums)) // Output: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
// }
