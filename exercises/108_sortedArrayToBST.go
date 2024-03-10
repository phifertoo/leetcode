package exercises

import "fmt"

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2                        // find midpoint
	root := &TreeNode{Val: nums[mid]}           // make root he midpoint
	root.Left = SortedArrayToBST(nums[:mid])    // Elements to the left of mid
	root.Right = SortedArrayToBST(nums[mid+1:]) // Elements to the right of mid

	return root
}

func SortedArrayToBSTTester() bool {

	fmt.Printf("%+v\n", SortedArrayToBST([]int{-10, -3, 0, 5, 9})) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

	return true
}
