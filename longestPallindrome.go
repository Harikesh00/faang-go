package main

// 43
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, end int

	for i := 0; i < len(s); i++ {
		// For odd length palindromes
		len1 := expandAroundCenter(s, i, i)
		// For even length palindromes
		len2 := expandAroundCenter(s, i, i+1)

		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// Return the length of the palindrome
	return right - left - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func main() {
//     fmt.Println(longestPalindrome("babad")) // Output: "bab" or "aba"
//     fmt.Println(longestPalindrome("cbbd"))  // Output: "bb"
// }
