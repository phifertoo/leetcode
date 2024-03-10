package exercises

import "fmt"

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Custom Judge:

// The judge will test your solution with the following code:

// int[] nums = [...]; // Input array
// int[] expectedNums = [...]; // The expected answer with correct length

// int k = removeDuplicates(nums); // Calls your implementation

// assert k == expectedNums.length;
// for (int i = 0; i < k; i++) {
//     assert nums[i] == expectedNums[i];
// }
// If all assertions pass, then your solution will be accepted.

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
