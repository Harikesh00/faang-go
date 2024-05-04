package main

// 47
func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := 1 // 1 for positive, -1 for negative
	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// Convert consecutive digits to a number
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			// Apply the previous sign and add the number to the result
			result += sign * num
			num = 0 // Reset the number
			sign = 1
		} else if s[i] == '-' {
			// Apply the previous sign and subtract the number from the result
			result += sign * num
			num = 0 // Reset the number
			sign = -1
		} else if s[i] == '(' {
			// Push the current result and sign onto the stack
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0 // Reset the result
			sign = 1
		} else if s[i] == ')' {
			// Apply the previous result and sign from the stack
			result += sign * num
			num = 0                       // Reset the number
			result *= stack[len(stack)-1] // Apply the sign
			stack = stack[:len(stack)-1]  // Pop the sign
			result += stack[len(stack)-1] // Add the previous result
			stack = stack[:len(stack)-1]  // Pop the previous result
		}
	}

	// Add the last number to the result
	result += sign * num

	return result
}

// func main() {
// 	fmt.Println(calculate("1 + 1"))         // Output: 2
// 	fmt.Println(calculate(" 2-1 + 2 "))     // Output: 3
// 	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)")) // Output: 23
// }
