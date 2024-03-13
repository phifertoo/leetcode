package exercises

import "fmt"

func Factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * Factorial(x-1)
}

func FactorialTester() bool {
	fmt.Print(Factorial(2)) //true
	fmt.Print(Factorial(4)) //true

	return true
}
