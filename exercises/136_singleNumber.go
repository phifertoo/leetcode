package exercises

import (
	"fmt"
)

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

func SingleNumber(nums []int) int {
	numberMap := make(map[int]bool)

	for _, number := range nums {
		_, exists := numberMap[number]

		if !exists {
			numberMap[number] = false
		} else {
			numberMap[number] = true
		}
	}

	for key := range numberMap {
		if !numberMap[key] {
			return key
		}
	}

	return 0
}

func SingleNumberTester() bool {
	fmt.Print(SingleNumber([]int{2, 2, 1}))       //1
	fmt.Print(SingleNumber([]int{4, 1, 2, 1, 2})) //4
	fmt.Print(SingleNumber([]int{1}))             //1

	return true
}
