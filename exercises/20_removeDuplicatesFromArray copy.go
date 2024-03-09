package exercises

import "fmt"

func removeDuplicatesFromArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 1 // Pointer for the position of the next unique element
	for i := 1; i < len(nums); i++ {
		// starting at index = 1, if the current number doesn't equal the previous
		// number, set the next number in the array to the current number
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j // j is the number of unique elements
}

func RemoveDuplicatesFromArrayTester() bool {
	fmt.Print(removeDuplicatesFromArray([]int{1, 1, 2}))                      //2
	fmt.Print(removeDuplicatesFromArray([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) //5

	return true
}
