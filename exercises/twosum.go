package leetcode

func twoSums(nums []int, target int) []int {
	indices := make(map[int]int)
	// create a hash map
	for i, element := range nums {
		indices[element] = i
	}

	// fmt.Printf("%+v\n", indices)

	for i, element := range nums {
		// fmt.Print(target, element)
		_, exists := indices[target-element]
		if exists && indices[target-element] != i {
			return []int{indices[target-element], i}
		}
	}

	return []int{}
}
