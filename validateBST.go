package main

// 28
/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	// Check if the current node's value is within the valid range
	if root.Val <= min || root.Val >= max {
		return false
	}

	// Check the left subtree with updated max value
	// and the right subtree with updated min value
	return isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}

func main() {
    // Example usage:
    root := &TreeNode{Val: 2}
    root.Left = &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 3}

    fmt.Println(isValidBST(root)) // Output: true
}
*/
