package stack

import "errors"

//Stack ...
type Stack []interface{}

//Pop ...
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}

//Push ...
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

// Top ...
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

//Cap ...
func (stack Stack) Cap() int {
	return cap(stack)
}

//Len ...
func (stack Stack) Len() int {
	return len(stack)
}

//IsEmpty ...
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
