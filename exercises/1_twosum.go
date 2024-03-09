package exercises

import "fmt"

func TwoSums(nums []int, target int) []int {
	indices := make(map[int]int)
	// create a hash map
	for i, element := range nums {
		indices[element] = i
	}

	for i, element := range nums {
		_, exists := indices[target-element]
		if exists && indices[target-element] != i {
			return []int{indices[target-element], i}
		}
	}

	return []int{}
}

func TwoSumsTester() bool {
	fmt.Print(TwoSums([]int{3, 2, 4}, 6)) //[0,1]

	return true
}
