package main

type StackInt struct {
	s []int
}

// isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
	return len(s.s) == 0
}

// length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
	return len(s.s)
}
