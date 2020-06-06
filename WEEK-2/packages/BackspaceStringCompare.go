package packages

import "fmt"

func isEqual(a, b []rune) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func backspaceCompare(S string, T string) bool {
	stackS := []rune{}
	stackT := []rune{}

	for _, chr := range S {
		if string(chr) == "#" {
			if len(stackS) != 0 {
				stackS = stackS[:len(stackS)-1]
			}
		} else {
			stackS = append(stackS, chr)
		}
	}

	for _, chr := range T {
		if string(chr) == "#" {
			if len(stackT) != 0 {
				stackT = stackT[:len(stackT)-1]
			}
		} else {
			stackT = append(stackT, chr)
		}
	}

	return isEqual(stackS, stackT)
}

// TestBackspaceCompare is the test function for backspace string compare
func TestBackspaceCompare() {
	fmt.Print("Testing for Backspace String Compare: ")
	fmt.Print(backspaceCompare("ab#c", "ad#c") == true, " ")
	fmt.Println(backspaceCompare("ab##", "c##d") == false)
}

// Problem Statement

// Given two strings S and T, return if they are equal when both are typed into empty text editors.
// # means a backspace character.

// Note that after backspacing an empty text, the text will continue empty.

// Example 1:

// Input: S = "ab#c", T = "ad#c"
// Output: true
// Explanation: Both S and T become "ac".

// Example 2:

// Input: S = "ab##", T = "c#d#"
// Output: true
// Explanation: Both S and T become "".

// Example 3:

// Input: S = "a##c", T = "#a#c"
// Output: true
// Explanation: Both S and T become "c".

// Example 4:

// Input: S = "a#c", T = "b"
// Output: false
// Explanation: S becomes "c" while T becomes "b".

// Note:

//     1 <= S.length <= 200
//     1 <= T.length <= 200
//     S and T only contain lowercase letters and '#' characters.

// Follow up:

//     Can you solve it in O(N) time and O(1) space?
