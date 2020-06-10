package packages

import "fmt"

func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	maxLen := 0

	m := make(map[int]int)
	m[0] = -1

	for i := range nums {
		if nums[i] == 0 {
			count--
		} else {
			count++
		}

		if _, check := m[count]; check {
			maxLen = max(maxLen, i-m[count])
		} else {
			m[count] = i
		}
	}

	return maxLen
}

// TestContigiousArray testing func for Contagious array
func TestContigiousArray() {
	fmt.Printf("Testing for Contagious Array: %t %t", findMaxLength([]int{0, 1}) == 2, findMaxLength([]int{0, 1, 0}) == 2)
}

// Problem Statement

// Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

// Example 1:

// Input: [0,1]
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.

// Example 2:

// Input: [0,1,0]
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

// Note: The length of the given binary array will not exceed 50,000.
