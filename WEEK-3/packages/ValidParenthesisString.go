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
