package exercises

import "fmt"

func removeElementFromArray(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0 // Pointer for the position of the next non-val element
	for i := 0; i < len(nums); i++ {
		// starting at index = 1, if the current number doesn't equal the previous
		// number, set the next number in the array to the current number
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j // j is the number of non-val elements
}

func RemoveElementFromArrayTester() bool {
	fmt.Print(removeElementFromArray([]int{3, 2, 2, 3}, 3))             //2
	fmt.Print(removeElementFromArray([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)) //5

	return true
}
