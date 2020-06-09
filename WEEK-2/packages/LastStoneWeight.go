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
