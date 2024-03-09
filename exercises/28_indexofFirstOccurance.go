package exercises

import (
	"fmt"
)

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

func indexOfFirstOccurance(haystack string, needle string) int {
	// return strings.Index(haystack, needle) // simple solution

	if len(needle) == 0 {
		return 0
	}

	// Only proceed if haystack is long enough to contain needle
	if len(haystack) >= len(needle) {
		// Iterate through haystack
		for i := 0; i <= len(haystack)-len(needle); i++ {
			// Check for needle starting at position i
			if haystack[i:i+len(needle)] == needle {
				return i // Found needle, return start index
			}
		}
	}

	// Needle not found
	return -1
}

func IndexOfFirstOccuranceTester() bool {
	fmt.Print(indexOfFirstOccurance("sadbutsad", "sad"))  //0
	fmt.Print(indexOfFirstOccurance("leetcode", "leeto")) //-1

	return true
}
