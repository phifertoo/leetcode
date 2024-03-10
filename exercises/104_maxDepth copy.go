package exercises

import "fmt"

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	} else {
		leftDepth := MaxDepth(root.Left)
		rightDepth := MaxDepth(root.Right)
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}

func MaxDepthTester() bool {
	node5 := TreeNode{Val: 7, Left: nil, Right: nil}
	node4 := TreeNode{Val: 15, Left: nil, Right: nil}
	node3 := TreeNode{Val: 20, Left: &node4, Right: &node5}
	node2 := TreeNode{Val: 9, Left: nil, Right: nil}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	fmt.Print(MaxDepth(&node1)) //3

	return true
}
