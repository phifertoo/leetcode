package exercises

import (
	"fmt"
	"regexp"
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

func IsPalindrome2(s string) bool {
	// Remove non-letter characters
	pattern := "[^A-Za-z0-9]"
	re1 := regexp.MustCompile(pattern)

	cleanedString := re1.ReplaceAllString(s, "")
	lowercasedString := strings.ToLower(cleanedString)

	for i, letter := range lowercasedString {
		if string(letter) != string(lowercasedString[len(lowercasedString)-1-i]) {
			return false
		}
	}

	return true
}

func IsPalindrome2Tester() bool {
	fmt.Print(IsPalindrome2("A man, a plan, a canal: Panama")) //true
	fmt.Print(IsPalindrome2("race a car"))                     //false
	fmt.Print(IsPalindrome2(" "))                              //true
	fmt.Print(IsPalindrome2("0P"))                             //false

	return true
}
