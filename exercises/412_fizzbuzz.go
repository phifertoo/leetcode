package exercises

import (
	"fmt"
	"strconv"
)

// Given an integer n, return a string array answer (1-indexed) where:

// answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
// answer[i] == "Fizz" if i is divisible by 3.
// answer[i] == "Buzz" if i is divisible by 5.
// answer[i] == i (as a string) if none of the above conditions are true.

func FizzBuzz(n int) []string {
	output := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, "FizzBuzz")
		} else if i%3 == 0 {
			output = append(output, "Fizz")
		} else if i%5 == 0 {
			output = append(output, "Buzz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}

	return output
}

func FizzBuzzTester() bool {
	fmt.Print(FizzBuzz(3))  //["1","2","Fizz"]
	fmt.Print(FizzBuzz(5))  //["1","2","Fizz","4","Buzz"]
	fmt.Print(FizzBuzz(15)) //["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]

	return true
}
