package exercises

import (
	"fmt"
)

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i, element := range nums {
		if element == target || target < element {
			if i == 0 {
				return 0
			}
			return i
		}
	}

	return len(nums)
}

func SearchInsertTester() bool {
	fmt.Print(SearchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Print(SearchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Print(SearchInsert([]int{1, 3, 5, 6}, 7)) // 4

	return true
}
