package main

// 35
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// If the current node is null or equals one of the nodes p or q, return the current node
	if root == nil || root == p || root == q {
		return root
	}

	// Recursively search for the LCA in the left and right subtrees
	leftLCA := lowestCommonAncestor(root.Left, p, q)
	rightLCA := lowestCommonAncestor(root.Right, p, q)

	// If one node is found in the left subtree and the other in the right subtree,
	// then the current node is the LCA
	if leftLCA != nil && rightLCA != nil {
		return root
	}

	// If both nodes are found in the left subtree, return the LCA from the left subtree
	if leftLCA != nil {
		return leftLCA
	}

	// If both nodes are found in the right subtree, return the LCA from the right subtree
	return rightLCA
}

// func main() {
//     // Example usage:
//     // Construct the binary tree:
//     //        3
//     //       / \
//     //      5   1
//     //     / \ / \
//     //    6  2 0  8
//     //      / \
//     //     7   4
//     root := &TreeNode{Val: 3}
//     root.Left = &TreeNode{Val: 5}
//     root.Right = &TreeNode{Val: 1}
//     root.Left.Left = &TreeNode{Val: 6}
//     root.Left.Right = &TreeNode{Val: 2}
//     root.Right.Left = &TreeNode{Val: 0}
//     root.Right.Right = &TreeNode{Val: 8}
//     root.Left.Right.Left = &TreeNode{Val: 7}
//     root.Left.Right.Right = &TreeNode{Val: 4}

//     p := root.Left
//     q := root.Right

//     lca := lowestCommonAncestor(root, p, q)
//     if lca != nil {
//         println("Lowest Common Ancestor:", lca.Val) // Output: 3
//     } else {
//         println("One or both of the nodes are not found in the tree.")
//     }
// }
