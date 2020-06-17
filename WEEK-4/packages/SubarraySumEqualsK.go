package packages

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0

	for i := range nums {
		sum := 0
		for _, val := range nums[i:] {
			sum += val
			if sum == k {
				count++
			}
		}
	}

	return count
}

// TestSubarraySum is the testing functon
func TestSubarraySum() {
	fmt.Printf("Testing for Min Path Sum: %t %t \n", subarraySum([]int{1, 1, 1}, 2) == 2, subarraySum([]int{1, 2, 3}, 3) == 2)
}
