package exercises

import "fmt"

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

func FirstUniqueCharacter(s string) int {
	letterMap := make(map[rune]bool)
	for _, letter := range s {
		_, exists := letterMap[letter]
		if exists {
			letterMap[letter] = true
		} else {
			letterMap[letter] = false
		}
	}

	for i, letter := range s {
		_, exists := letterMap[letter]
		if exists && !letterMap[letter] {
			return i
		}
	}

	return -1
}

func FirstUniqueCharacterTester() bool {
	fmt.Print(FirstUniqueCharacter("leetcode"))     //0
	fmt.Print(FirstUniqueCharacter("loveleetcode")) //2
	fmt.Print(FirstUniqueCharacter("aabb"))         //-1

	return true
}
