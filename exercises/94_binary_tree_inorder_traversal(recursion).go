package exercises

import "fmt"

// Given the root of a binary tree, return the inorder traversal of its nodes' values.
// in order means all the left nodes are visited 1st, current node 2nd, and right node 3rd

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	output := []int{}

	inorder(root, &output)

	return output
}

func inorder(node *TreeNode, output *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, output)
	// remember to use the reference because you are not returning anything
	*output = append(*output, node.Val)
	inorder(node.Right, output)
}

// func InorderTraversal(root *TreeNode) []int {
// 	output := []int{}

// 	inorder(root, &output)

// 	return output
// }

// func inorder(node *TreeNode, result *[]int) {
// 	if node == nil {
// 		return
// 	}
// 	// Traverse the left subtree
// 	inorder(node.Left, result)
// 	// append the root value of the current node
// 	*result = append(*result, node.Val)
// 	// Traverse the right subtree
// 	inorder(node.Right, result)
// }

func InorderTraversalTester() bool {
	node3 := TreeNode{Val: 3, Right: nil, Left: nil}
	node2 := TreeNode{Val: 2, Right: nil, Left: &node3}
	node1 := TreeNode{Val: 1, Right: &node2, Left: nil}
	fmt.Print(InorderTraversal(&node1)) // [1, 3,2]

	return true
}
