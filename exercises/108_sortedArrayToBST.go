package exercises

import "fmt"

// Given an integer array nums where the elements are sorted in ascending order, convert it to a
// height-balanced binary search tree.

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

	fmt.Printf("%+v\n", SortedArrayToBST([]int{-10, -3, 0, 5, 9})) //
	return true
}
