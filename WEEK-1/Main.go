package main

import (
	"fmt"
	"leetcode-30-days-challenge/WEEK-1/packages"
)

func main() {
	// Single Number
	fmt.Print("Testing for Single Number: ")
	fmt.Print(packages.SingleNumber([]int{2, 2, 1}) == 1, " ")
	fmt.Println(packages.SingleNumber([]int{4, 1, 2, 1, 2}) == 4)

	// Happy Number
	fmt.Print("Testing for Happy Number: ")
	fmt.Print(packages.IsHappy(19) == true, " ")
	fmt.Println(packages.IsHappy(2) == false)
}
