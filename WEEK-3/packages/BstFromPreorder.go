package packages

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func bstFromPreorder(preorder []int) *treeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &treeNode{Val: preorder[0]}
	}

	rIdx := len(preorder)
	rootVal := preorder[0]

	for i, val := range preorder {
		if val > rootVal {
			rIdx = i
		}
	}

	left := bstFromPreorder(preorder[1:rIdx])
	right := bstFromPreorder(preorder[rIdx:])

	return &treeNode{Val: rootVal, Left: left, Right: right}
}

func preorderTraversal(root *treeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}

	res = append(res, root.Val)

	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}

// TestBstFromPreorder is the testing func
func TestBstFromPreorder() {
	ar1 := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	ar2 := bstFromPreorder([]int{10, 5, 1, 7, 40, 50})

	test1 := fmt.Sprint(preorderTraversal(ar1))
	test2 := fmt.Sprint(preorderTraversal(ar2))

	fmt.Printf("Testing for BST from Preorder: %t %t \n", test1 == "[8 5 1 7 10 12]", test2 == "[10 5 1 7 40 50]")
}

// Problem Statement

// Return the root node of a binary search tree that matches the given preorder traversal.

// (Recall that a binary search tree is a binary tree where for every node, any descendant of
// node.left has a value < node.val, and any descendant of node.right has a value > node.val.
// Also recall that a preorder traversal displays the value of the node first, then traverses
// node.left, then traverses node.right.)

// It's guaranteed that for the given test cases there is always possible to find a binary search
// tree with the given requirements.

// Example 1:

// Input: [8,5,1,7,10,12]
// Output: [8,5,10,1,7,null,12]
