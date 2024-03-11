package exercises

import "fmt"

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

func MissingNumber(nums []int) int {
	numberMap := make(map[int]bool)

	for _, num := range nums {
		numberMap[num] = true
	}

	for i := 0; i < len(nums); i++ {
		_, exists := numberMap[i]
		if !exists {
			return i
		}
	}
	return len(nums)
}

func MissingNumberTester() bool {
	fmt.Print(MissingNumber([]int{3, 0, 1}))                   //2
	fmt.Print(MissingNumber([]int{0, 1}))                      //2
	fmt.Print(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) //8

	return true
}
