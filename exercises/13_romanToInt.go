package exercises

import (
	"fmt"
	"strings"
)

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.

func RomanToInt(s string) int {
	// replace all instances of IV and IX with 4 and 9
	// if IV or IX -2
	// if XL or XC -20
	// if CD or CM -200
	sum := 0
	if strings.Contains(s, "IV") || strings.Contains(s, "IX") {
		sum -= 2
	}

	if strings.Contains(s, "XL") || strings.Contains(s, "XC") {
		sum -= 20
	}

	if strings.Contains(s, "CD") || strings.Contains(s, "CM") {
		sum -= 200
	}

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		sum += romanMap[string(s[i])]
	}

	return sum
}

func RomanToIntTester() bool {
	fmt.Print(RomanToInt("III"))     //3
	fmt.Print(RomanToInt("LVIII"))   //58
	fmt.Print(RomanToInt("MCMXCIV")) //1994

	return true
}
