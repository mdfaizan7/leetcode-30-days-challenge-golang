package packages

import "fmt"

func checkVisits(grid *[][]byte, row, col int) {
	if row < 0 || row > len(*grid)-1 || col < 0 || col > len((*grid)[row])-1 || (*grid)[row][col] != '1' {
		return
	}

	(*grid)[row][col] = '2'

	checkVisits(grid, row+1, col)
	checkVisits(grid, row-1, col)
	checkVisits(grid, row, col+1)
	checkVisits(grid, row, col-1)
}

func numIslands(grid [][]byte) int {
	count := 0

	for row, rowEle := range grid {
		for col, colEle := range rowEle {
			if colEle != '1' {
				continue
			}

			checkVisits(&grid, row, col)
			count++
		}
	}

	return count
}

// TestNumIslands is the testing func for numIslands
func TestNumIslands() {
	arr1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}

	arr2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	fmt.Printf("Testing for Number of Islands: %t %t", numIslands(arr1) == 1, numIslands(arr2) == 3)
}

// Problem Statement

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.

// Example 1:

// Input:
// 11110
// 11010
// 11000
// 00000

// Output: 1

// Example 2:

// Input:
// 11000
// 11000
// 00100
// 00011

// Output: 3
