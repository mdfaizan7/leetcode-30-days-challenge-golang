package packages

// SingleNumber is the main function for this problem
func SingleNumber(nums []int) int {
	a := 0
	for i := range nums {
		a ^= nums[i]
	}

	return a
}

// Problem Statement

// Given a non-empty array of integers, every element appears twice except for one.
// Find that single one.

// Note:

// Your algorithm should have a linear runtime complexity. Could you implement it without
// using extra memory?

// Example 1:

// Input: [2,2,1]
// Output: 1

// Example 2:

// Input: [4,1,2,1,2]
// Output: 4
