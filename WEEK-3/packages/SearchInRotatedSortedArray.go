package packages

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	start := l
	l, r = 0, len(nums)-1

	if target >= nums[start] && target <= nums[r] {
		l = start
	} else {
		r = start
	}

	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

// TestRotatedArraySearch is the testing func
func TestRotatedArraySearch() {
	t := []int{4, 5, 6, 7, 0, 1, 2}

	fmt.Printf("Testing for Search in Rotated Sorted Array: %t %t \n", search(t, 0) == 4, search(t, 3) == -1)
}
