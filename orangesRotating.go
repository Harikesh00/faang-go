package main

// 30
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	freshOranges := 0
	queue := make([][2]int, 0)

	// Initialize the queue with the coordinates of all initially rotten oranges
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	minutes := 0
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[i]
			for _, dir := range directions {
				newX, newY := curr[0]+dir[0], curr[1]+dir[1]
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					freshOranges--
					queue = append(queue, [2]int{newX, newY})
				}
			}
		}

		queue = queue[size:]
		if len(queue) > 0 {
			minutes++
		}
	}

	if freshOranges == 0 {
		return minutes
	} else {
		return -1
	}
}

// func main() {
//     grid := [][]int{
//         {2, 1, 1},
//         {1, 1, 0},
//         {0, 1, 1},
//     }

//     fmt.Println(orangesRotting(grid)) // Output: 4
// }
