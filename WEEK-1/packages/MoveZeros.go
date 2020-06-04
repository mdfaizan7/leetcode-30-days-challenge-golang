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
