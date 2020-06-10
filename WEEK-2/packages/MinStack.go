package packages

import "fmt"

// MinStack is definition for struct minnStack
type MinStack struct {
	Stack []int
	Min   []int
}

// Constructor for making a MinStack
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}
}

func (stack *MinStack) push(x int) {
	stack.Stack = append(stack.Stack, x)
	if len(stack.Min) == 0 {
		stack.Min = append(stack.Min, x)
	} else {
		min := stack.Min[len(stack.Min)-1]
		if min >= x {
			stack.Min = append(stack.Min, x)
		}
	}
}

func (stack *MinStack) pop() {
	x := stack.Stack[len(stack.Stack)-1]
	stack.Stack = stack.Stack[:len(stack.Stack)-1]

	min := stack.Min[len(stack.Min)-1]

	if min == x {
		stack.Min = stack.Min[:len(stack.Min)-1]
	}
}

func (stack *MinStack) top() int {
	return stack.Stack[len(stack.Stack)-1]
}

func (stack *MinStack) getMin() int {
	return stack.Min[len(stack.Min)-1]
}

// TestMinStack is for testing the MinStack
func TestMinStack() {
	minStack := Constructor()
	minStack.push(-2)
	minStack.push(0)
	minStack.push(-3)
	fmt.Print("Testing for Min Stack: ")
	fmt.Print(minStack.getMin() == -3, " ")
	minStack.pop()
	fmt.Println(minStack.top() == 0 && minStack.getMin() == -2)
}

// Problem Statement

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

//     push(x) -- Push element x onto stack.
//     pop() -- Removes the element on top of the stack.
//     top() -- Get the top element.
//     getMin() -- Retrieve the minimum element in the stack.

// Example 1:

// Input
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// Output
// [null,null,null,null,-3,null,0,-2]

// Explanation
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -2

// Constraints:
//     Methods pop, top and getMin operations will always be called on non-empty stacks.
