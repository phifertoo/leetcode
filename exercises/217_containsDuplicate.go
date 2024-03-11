package exercises

import "fmt"

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

func ContainsDuplicate(nums []int) bool {
	foundNumbers := make(map[int]bool)

	for _, number := range nums {
		_, exists := foundNumbers[number]
		if exists {
			return true
		} else {
			foundNumbers[number] = true
		}
	}

	return false
}

func ContainsDuplicateTester() bool {
	fmt.Print(ContainsDuplicate([]int{1, 2, 3, 1}))                   //true
	fmt.Print(ContainsDuplicate([]int{1, 2, 3, 4}))                   //false
	fmt.Print(ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) //true

	return true
}
