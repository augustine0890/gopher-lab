package main

import "fmt"

type StackInt struct {
	items []int
}

// isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
	return len(s.items) == 0
}

// length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
	return len(s.items)
}

// The print function will print the elements of the array.
func (s *StackInt) Print() {
	fmt.Println(s.items)
}

// push() function append value to the data.
func (s *StackInt) Push(value int) {
	s.items = append(s.items, value)
}

/* In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it. */

func (s *StackInt) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

/*
top() function will first check that the stack is not empty
then returns the value stored in the top element
of stack (does not remove it).
*/
func (s *StackInt) Top() int {
	if s.IsEmpty() {
		return 0
	}
	item := s.items[len(s.items)-1]
	return item
}
