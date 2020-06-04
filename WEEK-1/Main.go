package main

import (
	"fmt"
	"leetcode-30-days-challenge/WEEK-1/packages"
)

func testEq(a, b []int) bool {
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

func main() {
	// Single Number
	fmt.Print("Testing for Single Number: ")
	fmt.Print(packages.SingleNumber([]int{2, 2, 1}) == 1, " ")
	fmt.Println(packages.SingleNumber([]int{4, 1, 2, 1, 2}) == 4)

	// Happy Number
	fmt.Print("Testing for Happy Number: ")
	fmt.Print(packages.IsHappy(19) == true, " ")
	fmt.Println(packages.IsHappy(2) == false)

	// Maximum Subarray
	fmt.Print("Testing for Maximum Subarray: ")
	fmt.Print(packages.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6, " ")
	fmt.Println(packages.MaxSubArray([]int{-2, 1, -5, 4}) == 4)

	// Move Zeros
	fmt.Print("Testing for Move Zeros: ")
	fmt.Print(testEq(packages.MoveZeroes([]int{0, 0, 1}), []int{1, 0, 0}), " ")
	fmt.Println(testEq(packages.MoveZeroes([]int{0, 1, 0, 3, 12}), []int{1, 3, 12, 0, 0}))

	// Best Time to Buy and Sell Stock II
	fmt.Print("Testing for Best Time to Buy and Sell Stock II: ")
	fmt.Print(packages.MaxProfit([]int{7, 1, 5, 3, 6, 4}) == 7, " ")
	fmt.Println(packages.MaxProfit([]int{1, 2, 3, 4, 5}) == 4)
}
