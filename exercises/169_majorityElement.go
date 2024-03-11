package exercises

import "fmt"

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

func MajorityElement(nums []int) int {
	numberMap := make(map[int]int)

	for _, num := range nums {
		_, exists := numberMap[num]
		if exists {
			numberMap[num]++
		} else {
			numberMap[num] = 1
		}

		if numberMap[num] > (len(nums) / 2) {
			return num
		}
	}

	// fmt.Printf("%+v\n", numberMap)

	return 0
}

func MajorityElementTester() bool {
	fmt.Print(MajorityElement([]int{3, 2, 3}))             // 3
	fmt.Print(MajorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2

	return true
}
