package exercises

import (
	"fmt"
	"strconv"
)

// Given an integer x, return true if x is a palindrome, and false otherwise.

func ReserseBits(num uint32) uint32 {
	binaryStr := strconv.FormatUint(uint64(num), 2)

	// Pad the string with leading zeros to ensure it has a length of 32
	for len(binaryStr) < 32 {
		binaryStr = "0" + binaryStr
	}

	// Reverse the string
	reversedStr := ""
	for _, bit := range binaryStr {
		reversedStr = string(bit) + reversedStr
	}

	// Convert the reversed binary string back to a number
	result, err := strconv.ParseUint(reversedStr, 2, 32)
	if err != nil {
		fmt.Println("Error converting binary string to number:", err)
		return 0
	}

	return uint32(result)
}

func ReserseBitsTester() bool {
	// 00000010100101000001111010011100
	fmt.Print(ReserseBits(uint32(43261596))) //964176192
	// 11111111111111111111111111111101
	fmt.Print(ReserseBits(uint32(4294967293))) //3221225471

	return true
}
