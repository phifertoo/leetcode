package exercises

import "fmt"

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	uncommonIndex := getIndex(strs)
	if uncommonIndex == 0 {
		return ""
	}

	return string(strs[0][0:uncommonIndex])
}

func getIndex(strs []string) int {
	firstElement := strs[0]

	// iterate through each letter of first element across all words in array to see where the first letter is not found
	uncommonIndex := 0

	// if this for loop finishes, all the letters are the same so we should return the length of the first element
	for i := 0; i < len(firstElement); i++ {
		for j := 0; j < len(strs); j++ {
			if i > len(strs[j])-1 || firstElement[i] != strs[j][i] {
				uncommonIndex = i
				return uncommonIndex
			}
		}
	}

	return len(firstElement)
}

func LongestCommonPrefixTester() bool {
	fmt.Print(longestCommonPrefix([]string{"flower", "flow", "flight"}))   // "fl"
	fmt.Print(longestCommonPrefix([]string{"dog", "racecar", "car"}))      // ""
	fmt.Print(longestCommonPrefix([]string{"ab", "a"}))                    // ""
	fmt.Print(longestCommonPrefix([]string{"flower", "flower", "flower"})) // ""

	return true
}
