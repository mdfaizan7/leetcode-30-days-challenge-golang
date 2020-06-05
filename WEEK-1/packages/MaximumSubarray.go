package packages

// MaxSubArray is the main func for this problem
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currSum := nums[0]

	for _, val := range nums[1:] {
		currSum = max(currSum+val, val)
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Problem Statement

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

// Follow up:

// If you have figured out the O(n) solution, try coding another solution using the divide and conquer
// approach, which is more subtle.
