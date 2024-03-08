package exercises

import (
	"fmt"
	"strings"
)

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
