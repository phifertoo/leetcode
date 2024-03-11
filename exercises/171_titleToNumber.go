package exercises

import "fmt"

// Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

func TitleToNumber(columnTitle string) int {
	letterMap := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5,
		"F": 6, "G": 7, "H": 8, "I": 9, "J": 10,
		"K": 11, "L": 12, "M": 13, "N": 14, "O": 15,
		"P": 16, "Q": 17, "R": 18, "S": 19, "T": 20,
		"U": 21, "V": 22, "W": 23, "X": 24, "Y": 25,
		"Z": 26,
	}

	// keep a running sum
	// multiply the previous total by the total then add the current letter value
	total := 0
	for _, letter := range columnTitle {
		total = total*26 + letterMap[string(letter)]
	}

	return total
}

func TitleToNumberTester() bool {
	fmt.Print(TitleToNumber("A"))  // 1
	fmt.Print(TitleToNumber("AB")) // 28
	fmt.Print(TitleToNumber("ZY")) // 701

	return true
}

// length * row quantity - 1

// Length + 1 squared
// aa5, ab6, ac7, ad 8
// ba7, bb8, bc9
// ca10, ca11, ca12

// 3 * (3 - 1)
// aaa13
