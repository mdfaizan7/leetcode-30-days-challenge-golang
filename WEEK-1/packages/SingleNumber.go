package packages

// SingleNumber is the main function for this problem
func SingleNumber(nums []int) int {
	a := 0
	for i := range nums {
		a ^= nums[i]
	}

	return a
}
