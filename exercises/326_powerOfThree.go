package exercises

import "fmt"

// Given an integer n, return true if it is a power of three. Otherwise, return false.

// An integer n is a power of three, if there exists an integer x such that n == 3x.

func PowerOfThree(n int) bool {
	return dividByThree(n)
}

func dividByThree(n int) bool {

	if n == 1 {
		return true
	}

	quotient := n / 3

	if n%3 != 0 || n == 0 {
		return false
	}

	return dividByThree(quotient)
}

func PowerOfThreeTester() bool {
	// fmt.Print(PowerOfThree(27)) //true
	// fmt.Print(PowerOfThree(0)) //false
	fmt.Print(PowerOfThree(1)) //true
	// fmt.Print(PowerOfThree(45)) //false
	// fmt.Print(PowerOfThree(-1)) //false

	return true
}
