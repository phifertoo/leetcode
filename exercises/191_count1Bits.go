package exercises

import (
	"fmt"
	"strconv"
)

// Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

func Count1Bits(num uint32) int {
	// convert to binary string
	binaryStr := strconv.FormatUint(uint64(num), 2)

	fmt.Print(binaryStr)
	count := 0
	for _, number := range binaryStr {
		if number == '1' {
			count++
		}
	}

	return count

}

func Count1BitsTester() bool {
	// 00000000000000000000000000001011
	// fmt.Print(Count1Bits(uint32(11))) //3
	// 00000000000000000000000010000000
	fmt.Print(Count1Bits(uint32(128))) //1

	return true
}
