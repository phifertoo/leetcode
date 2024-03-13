package exercises

import "fmt"

func Fibonacci(x int) int {
	if x < 2 {
		return x
	}

	// 7 => 6, 5
	// 6 => 5, 4

	return Fibonacci(x-1) + Fibonacci(x-2)
}

func FibonacciTester() bool {
	fmt.Print(Fibonacci(7)) //13

	return true
}
