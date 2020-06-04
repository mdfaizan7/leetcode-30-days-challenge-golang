package main

import "fmt"

func singleNumber(nums []int) int {
	a := 0
	for i := range nums {
		a ^= nums[i]
	}

	return a
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}) == 1)
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}) == 4)
}
