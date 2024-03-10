package exercises

import "fmt"

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	return leftNode.Val == rightNode.Val && isMirror(leftNode.Right, rightNode.Left) && isMirror(rightNode.Right, leftNode.Left)
}

func IsSymmetricTester1() bool {
	node7 := TreeNode{Val: 3, Left: nil, Right: nil}
	node6 := TreeNode{Val: 4, Left: nil, Right: nil}
	node5 := TreeNode{Val: 4, Left: nil, Right: nil}
	node4 := TreeNode{Val: 3, Left: nil, Right: nil}
	node3 := TreeNode{Val: 2, Left: &node6, Right: &node7}
	node2 := TreeNode{Val: 2, Left: &node4, Right: &node5}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	fmt.Print(IsSymmetric(&node1)) //true

	return true
}

func IsSymmetricTester2() bool {
	node5 := TreeNode{Val: 3, Left: nil, Right: nil}
	node4 := TreeNode{Val: 3, Left: nil, Right: nil}
	node3 := TreeNode{Val: 2, Left: nil, Right: &node5}
	node2 := TreeNode{Val: 2, Left: nil, Right: &node4}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	fmt.Print(IsSymmetric(&node1)) //false

	return true
}
