package main

// 42
/*
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    var result []int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        // Add the rightmost node's value to the result
        result = append(result, queue[len(queue)-1].Val)

        // Process the current level
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            node := queue[i]
            // Enqueue the child nodes
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        // Remove processed nodes from the queue
        queue = queue[levelSize:]
    }

    return result
}

func main() {
    // Example usage:
    // Construct the binary tree:
    //     1
    //    / \
    //   2   3
    //    \   \
    //     5   4
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 5}
    root.Right.Right = &TreeNode{Val: 4}

    fmt.Println(rightSideView(root)) // Output: [1 3 4]
}
*/
