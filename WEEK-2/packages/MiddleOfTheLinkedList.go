package packages

import "fmt"

// MyLinkedList is the definition for singly-linked list
type MyLinkedList struct {
	Head *ListNode
}

// ListNode is the definition for the node of a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ll *MyLinkedList) pushFront(Val int) *ListNode {
	node := &ListNode{Val, nil}
	if ll.Head == nil {
		ll.Head = node
		return ll.Head
	}
	node.Next = ll.Head
	ll.Head = node

	return ll.Head
}

func print(ll MyLinkedList) string {
	currentNode := ll.Head
	linkedList := ""
	for currentNode != nil {
		linkedList += fmt.Sprintf("%d ", currentNode.Val)
		currentNode = currentNode.Next
	}

	return linkedList
}

func middleNode(head *ListNode) *ListNode {
	i, j := head, head

	for j != nil && j.Next != nil {
		j = j.Next.Next
		i = i.Next
	}

	return i
}

// TestMiddleNode is the test function for checking the middle Node
func TestMiddleNode() {
	var test1 = MyLinkedList{}
	var test2 = MyLinkedList{}

	for i := 4; i > -1; i-- {
		test1.pushFront(i)
	}
	for i := 99; i > -1; i -= 2 {
		test2.pushFront(i)
	}

	test1.Head = middleNode(test1.Head)
	test2.Head = middleNode(test2.Head)

	fmt.Print("Testing for Middle of the Linked List: ")
	fmt.Print(print(test1) == "2 3 4 ", " ")
	fmt.Println(print(test2) == "51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95 97 99 ")
}

// Problem Statement

// Given a non-empty, singly linked list with head node head, return a middle node of linked list.

// If there are two middle nodes, return the second middle node.

// Example 1:

// Input: [1,2,3,4,5]
// Output: Node 3 from this list (Serialization: [3,4,5])
// The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
// Note that we returned a ListNode object ans, such that:
// ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.

// Example 2:

// Input: [1,2,3,4,5,6]
// Output: Node 4 from this list (Serialization: [4,5,6])
// Since the list has two middle nodes with values 3 and 4, we return the second one.

// Note:
//     The number of nodes in the given list will be between 1 and 100
