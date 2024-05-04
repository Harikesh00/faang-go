package main

// 29
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	numIslands := 0
	rows, cols := len(grid), len(grid[0])

	// Define a helper function for DFS
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == '0' {
			return
		}

		// Mark the current cell as visited
		grid[row][col] = '0'

		// Visit neighboring cells
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	// Traverse the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				numIslands++
				dfs(i, j)
			}
		}
	}

	return numIslands
}

// func main() {
// 	grid := [][]byte{
// 		{'1', '1', '0', '0', '0'},
// 		{'1', '1', '0', '0', '0'},
// 		{'0', '0', '1', '0', '0'},
// 		{'0', '0', '0', '1', '1'},
// 	}

// 	fmt.Println(numIslands(grid)) // Output: 3
// }
