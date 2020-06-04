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
