package exercises

import "fmt"

// Given an integer x, return true if x is a palindrome, and false otherwise.

func IsPalindrome(x int) bool {
	// convert number to string
	numStr := fmt.Sprintf("%d", x)

	output := true
	for i := 0; i < len(numStr); i++ {
		if numStr[len(numStr)-1-i] != numStr[i] {
			return false
		}
	}

	return output
}

func IsPalindromeTester() bool {
	fmt.Print(IsPalindrome(5544)) //true
	fmt.Print(IsPalindrome(121))  //true
	fmt.Print(IsPalindrome(-121)) //false
	fmt.Print(IsPalindrome(10))   //false

	return true
}
