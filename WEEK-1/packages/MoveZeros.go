package packages

// MoveZeroes is the main func for this problem
func MoveZeroes(nums []int) []int {
	count := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
	}

	return nums
}

// Problem Statement

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative
// order of the non-zero elements.

// Example:

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Note:

//     You must do this in-place without making a copy of the array.
//     Minimize the total number of operations.
