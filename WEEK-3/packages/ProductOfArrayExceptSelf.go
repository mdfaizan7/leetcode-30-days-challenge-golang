package packages

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	ans[0] = 1

	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	r := 1

	for i := len(nums) - 1; i > -1; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}

	return ans
}

func isEqual(a, b []int) bool {
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

// TestProductExceptSlef is for testing this problem
func TestProductExceptSlef() {
	fmt.Printf("Testing for Product of Array Except Self: %t %t \n",
		isEqual(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}),
		isEqual(productExceptSelf([]int{12, 2, 3, 1, 4, 6, 3, 6}),
			[]int{2592, 15552, 10368, 31104, 7776, 5184, 10368, 5184}))
}

// Problem Statement

// Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal
// to the product of all the elements of nums except nums[i].

// Example:

// Input:  [1,2,3,4]
// Output: [24,12,8,6]

// Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the
// array (including the whole array) fits in a 32 bit integer.

// Note: Please solve it without division and in O(n).
