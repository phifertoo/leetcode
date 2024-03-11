package exercises

import "fmt"

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func IsValidAnagram(s string, t string) bool {
	wordMap := make(map[rune]int)

	for _, letter := range s {
		_, exists := wordMap[letter]
		if exists {
			wordMap[letter]++
		} else {
			wordMap[letter] = 1
		}
	}

	for _, letter := range t {
		_, exists := wordMap[letter]
		if exists {
			wordMap[letter]--
		} else {
			return false
		}
	}

	for key := range wordMap {
		if wordMap[key] != 0 {
			return false
		}
	}

	return true
}

func IsValidAnagramTester() bool {
	fmt.Print(IsValidAnagram("anagram", "nagaram")) //true
	fmt.Print(IsValidAnagram("rat", "car"))         //false

	return true
}
