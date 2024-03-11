package exercises

import "fmt"

// Write a function that reverses a string. The input string is given as an array of characters s.

// You must do this by modifying the input array in-place with O(1) extra memory.

func ReverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		// you have to do this at the same time!!!
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}

func ReverseStringTester() bool {
	fmt.Print(ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})) // olleh

	return true
}
