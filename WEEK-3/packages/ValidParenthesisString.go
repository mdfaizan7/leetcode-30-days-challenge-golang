package packages

import "fmt"

func checkValidString(s string) bool {
	stack := []int{}
	starStack := []int{}

	for i, char := range s {
		if string(char) == "(" {
			stack = append(stack, i)
		} else if string(char) == ")" {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else if len(starStack) != 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		} else {
			starStack = append(starStack, i)
		}
	}

	for len(stack) != 0 {
		if len(starStack) == 0 {
			return false
		} else if starStack[len(starStack)-1] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			starStack = starStack[:len(starStack)-1]
		} else {
			return false
		}
	}

	return true
}

// TestValidString is the testing func for this problem
func TestValidString() {
	fmt.Printf("Testing for Valid Parenthesis String: %t %t \n",
		checkValidString(")") == false, checkValidString("(*))") == true)
}

// Problem Statement

//  Given a string containing only three types of characters: '(', ')' and '*',
// write a function to check whether this string is valid. We define the validity of a string by these rules:

//     Any left parenthesis '(' must have a corresponding right parenthesis ')'.

//     Any right parenthesis ')' must have a corresponding left parenthesis '('.

//     Left parenthesis '(' must go before the corresponding right parenthesis ')'.

//     '*' could be treated as a single right parenthesis ')' or a single left parenthesis
//		 '(' or an empty string.

//     An empty string is also valid.

// Example 1:

// Input: "()"
// Output: True

// Example 2:

// Input: "(*)"
// Output: True

// Example 3:

// Input: "(*))"
// Output: True

// Note:
//     The string size will be in the range [1, 100].
