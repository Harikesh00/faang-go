package main

import (
	"unicode"
)

// 39
func myAtoi(s string) int {
	// Skip leading whitespaces
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check sign
	sign := 1
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert digits
	result := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')
		if result > (1<<31-1-digit)/10 {
			if sign == -1 {
				return -1 << 31
			} else {
				return 1<<31 - 1
			}
		}
		result = result*10 + digit
		i++
	}

	return result * sign
}

// func main() {
//     fmt.Println(myAtoi("42"))            // Output: 42
//     fmt.Println(myAtoi("   -42"))        // Output: -42
//     fmt.Println(myAtoi("4193 with words")) // Output: 4193
//     fmt.Println(myAtoi("words and 987"))  // Output: 0
//     fmt.Println(myAtoi("-91283472332"))   // Output: -2147483648 (minimum 32-bit integer)
// }
