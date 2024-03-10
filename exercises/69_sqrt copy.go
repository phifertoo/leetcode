package exercises

import "fmt"

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

// You must not use any built-in exponent function or operator.

func Sqrt(x int) int {
	i := 1

	for i*i <= x {
		i++
	}

	return i - 1
}

func SqrtTester() bool {
	fmt.Print(Sqrt(4)) //2
	fmt.Print(Sqrt(8)) //2

	return true
}
