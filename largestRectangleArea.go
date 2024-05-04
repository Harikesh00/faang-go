package main

// 50
/*
func largestRectangleArea(heights []int) int {
    n := len(heights)
    stack := make([]int, 0)
    maxArea := 0

    for i := 0; i <= n; i++ {
        var h int
        if i < n {
            h = heights[i]
        } else {
            h = 0 // Consider the last bar with height 0 to clear the stack
        }

        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            // Pop the top bar from the stack
            topIdx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            // Calculate the width of the rectangle
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            // Calculate the area of the rectangle formed by the popped bar
            area := widths[topIdx] * width

            // Update the maximum area found so far
            maxArea = max(maxArea, area)
        }

        // Push the current index onto the stack
        stack = append(stack, i)
    }

    return maxArea
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // Output: 10
}
*/
