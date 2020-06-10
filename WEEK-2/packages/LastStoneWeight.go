package packages

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		x, y := stones[len(stones)-1], stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if x != y {
			stones = append(stones, x-y)
		}

		sort.Ints(stones)
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}

// TestLastStoneWeight is the testib=ng func for this problem
func TestLastStoneWeight() {
	fmt.Print("Testing for Last Stone Weight: ")
	fmt.Print(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1, " ")
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 65, 12, 123, 12, 1}) == 11)
}

// Problem Statement

// We have a collection of stones, each stone has a positive integer weight.

// Each turn, we choose the two heaviest stones and smash them together.
// Suppose the stones have weights x and y with x <= y. The result of this smash is:

//     If x == y, both stones are totally destroyed;
//     If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.

// At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

// Example 1:

// Input: [2,7,4,1,8,1]
// Output: 1
// Explanation:
// We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
// we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
// we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
// we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.

// Note:

//     1 <= stones.length <= 30
//     1 <= stones[i] <= 1000
