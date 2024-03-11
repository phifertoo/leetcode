package exercises

import (
	"fmt"
	"strconv"
)

// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

// when thinking about loops, check if seen

func HappyNumber(n int) bool {

	targetNumber := n
	seenMap := make(map[int]bool)

	for targetNumber != 1 {

		numStr := strconv.Itoa(targetNumber)
		tempSum := 0
		for _, number := range numStr {
			parsedInt, err := strconv.Atoi(string(number))
			if err != nil {
				return false
			}
			tempSum += parsedInt * parsedInt
		}
		fmt.Print("tempSum", tempSum)

		_, exists := seenMap[tempSum]
		if exists {
			return false
		} else {
			seenMap[tempSum] = true
		}
		targetNumber = tempSum
	}

	return true
}

func HappyNumberTester() bool {
	fmt.Print(HappyNumber(19)) //true
	fmt.Print(HappyNumber(2))  //false
	fmt.Print(HappyNumber(3))  //false

	return true
}
