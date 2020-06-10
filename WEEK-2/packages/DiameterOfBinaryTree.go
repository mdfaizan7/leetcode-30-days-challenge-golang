package packages

import "fmt"

// TreeNode is the Definition of the Binary Tree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDepth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	leftDepth := getDepth(root.Left, diameter)
	rightDepth := getDepth(root.Right, diameter)

	*diameter = max(*diameter, leftDepth+rightDepth)

	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter := 0

	leftDepth := getDepth(root.Left, &diameter)
	rightDepth := getDepth(root.Right, &diameter)

	return max(leftDepth+rightDepth, diameter)
}

// ConstructTreeNode is the constructor for TreeNode
func ConstructTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func makeLevelorderBinaryTree(arr []int, root *TreeNode, i, n int) *TreeNode {
	if i < n {
		temp := ConstructTreeNode(arr[i])
		root = temp

		root.Left = makeLevelorderBinaryTree(arr, root.Left, 2*i+1, n)
		root.Right = makeLevelorderBinaryTree(arr, root.Right, 2*i+2, n)
	}

	return root
}

// TestDiameterOfBinaryTree is the test function for checking the diameter of a btree
func TestDiameterOfBinaryTree() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, -7, -3, 0, 0, -9, -3, 9, -7, -4, 0, 6, 0, -6, -6, 0, 0, 1, 6, 5, 0, 9, 0, 0, -1, -4, 0, 0, 0, -2} // consider 0 as nil
	bTree1 := makeLevelorderBinaryTree(arr1, nil, 0, len(arr1))
	bTree2 := makeLevelorderBinaryTree(arr2, nil, 0, len(arr2))

	fmt.Print("Testing for Diameter of Binary Tree: ")
	fmt.Print(diameterOfBinaryTree(bTree1) == 3, " ")
	fmt.Println(diameterOfBinaryTree(bTree2) == 8)
}

// Problem Statement

// Given a binary tree, you need to compute the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
// This path may or may not pass through the root.

// Example:
// Given a binary tree

//           1
//          / \
//         2   3
//        / \
//       4   5

// Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

// Note: The length of path between two nodes is represented by the number of edges between them.
