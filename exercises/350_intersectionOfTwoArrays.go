package exercises

import "fmt"

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

func IntersectionOfTwoArrays(nums1 []int, nums2 []int) []int {
	numberMap := make(map[int]int)

	short := nums1
	long := nums2

	if len(nums1) > len(nums2) {
		long = nums1
		short = nums2
	}

	for _, number := range long {
		_, exists := numberMap[number]
		if exists {
			numberMap[number]++
		} else {
			numberMap[number] = 1
		}
	}

	output := []int{}

	for _, number := range short {
		_, exists := numberMap[number]
		if exists && numberMap[number] > 0 {
			numberMap[number]--
			output = append(output, number)
		}
	}

	return output
}

func IntersectionOfTwoArraysTester() bool {
	fmt.Print(IntersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))       // [2, 2]
	fmt.Print(IntersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // [4,9]

	return true
}
