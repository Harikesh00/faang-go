package main

// 40
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var result []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse bottom row (if applicable)
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse left column (if applicable)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

// func main() {
//     matrix := [][]int{
//         {1, 2, 3},
//         {4, 5, 6},
//         {7, 8, 9},
//     }
//     fmt.Println(spiralOrder(matrix)) // Output: [1 2 3 6 9 8 7 4 5]
// }
