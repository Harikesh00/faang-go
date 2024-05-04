package main

// 38
func sortColors(nums []int) {
	// Initialize three pointers
	red, white, blue := 0, 0, len(nums)-1

	// Iterate through the array
	for white <= blue {
		switch nums[white] {
		case 0: // If the element is red (0), swap it with the element at the red pointer and move both pointers to the right
			nums[red], nums[white] = nums[white], nums[red]
			red++
			white++
		case 1: // If the element is white (1), just move the white pointer to the right
			white++
		case 2: // If the element is blue (2), swap it with the element at the blue pointer and move the blue pointer to the left
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}

// func main() {
//     nums := []int{2, 0, 2, 1, 1, 0}
//     sortColors(nums)
//     fmt.Println(nums) // Output: [0 0 1 1 2 2]
// }
