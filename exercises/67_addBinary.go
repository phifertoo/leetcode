package exercises

import (
	"fmt"
	"strings"
)

// Given two binary strings a and b, return their sum as a binary string.

func AddBinary(a string, b string) string {
	// determine longer string
	longer := ""
	shorter := ""
	if len(a) >= len(b) {
		longer = a
		shorter = b
	} else {
		longer = b
		shorter = a
	}

	// create output slice
	output := []string{}
	for i := 0; i < len(longer); i++ {
		output = append(output, string(longer[i]))
	}

	carryOver := false
	for i := len(longer) - 1; i >= 0; i-- {
		shortIndex := len(shorter) - 1 - (len(longer) - 1 - i)
		shortNum := "0"
		if shortIndex > -1 {
			shortNum = string(shorter[shortIndex])
		} else {
			shortNum = "0"
		}

		num, next := handleBinary(shortNum, string(longer[i]), carryOver)
		if next {
			carryOver = true
		} else {
			carryOver = false
		}

		output[i] = num
	}

	if carryOver {
		output = append([]string{"1"}, output...)
	}

	return strings.Join(output, "")
}

func handleBinary(numA string, numB string, carryOver bool) (string, bool) {
	if carryOver {
		if numA == "1" && numB == "1" {
			return "1", true
		}

		if numA == "1" || numB == "1" {
			return "0", true
		}

		return "1", false
	}

	if numA == "1" && numB == "1" {
		return "0", true
	}

	if numA == "1" || numB == "1" {
		return "1", false
	}

	return "0", false
}

func AddBinaryTester() bool {
	fmt.Print(AddBinary("11", "1"))      // "100"
	fmt.Print(AddBinary("1010", "1011")) //"10101"

	return true
}
