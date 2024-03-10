package exercises

import "fmt"

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

func PascalsTriangle(numRows int) [][]int {
	output := [][]int{}

	if numRows == 0 {
		return output
	}

	output = append(output, []int{1})

	if numRows > 1 {
		// row 3 => i = 3
		for i := 1; i < numRows; i++ {
			innerSlice := []int{1}
			for j := 1; j <= i; j++ {
				if j == i {
					innerSlice = append(innerSlice, 1)
				} else {
					innerSlice = append(innerSlice, output[i-1][j-1]+output[i-1][j])
				}

			}
			output = append(output, innerSlice)
		}
	}

	return output
}

func PascalsTriangleTester() bool {

	fmt.Printf("%+v\n", PascalsTriangle(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

	return true
}
