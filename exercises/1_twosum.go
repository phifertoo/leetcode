package exercises

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

func TwoSums(nums []int, target int) []int {
	indices := make(map[int]int)
	// create a hash map
	for i, element := range nums {
		indices[target-element] = i
	}

	for i, element := range nums {
		_, exists := indices[element]
		if exists && element != target-element {
			return []int{i, indices[element]}
		}
	}

	return []int{}
}

// func TwoSums(nums []int, target int) []int {
// 	indices := make(map[int]int)
// 	// create a hash map
// 	for i, element := range nums {
// 		indices[element] = i
// 	}

// 	for i, element := range nums {
// 		_, exists := indices[target-element]
// 		if exists && indices[target-element] != i {
// 			return []int{indices[target-element], i}
// 		}
// 	}

// 	return []int{}
// }

func TwoSumsTester() bool {
	fmt.Print(TwoSums([]int{3, 2, 4}, 6)) //[0,1]

	return true
}
