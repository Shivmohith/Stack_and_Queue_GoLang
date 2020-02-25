package stack

import "errors"

// Structure which stores a slice array of type int32

type StackArray struct {
	Array []int32
}

// Function to create a new empty stack
func NewStackArray() *StackArray {
	return &StackArray{Array: make([]int32, 0)}
}

type StackMethods interface {
	Push(int32)
	Pop() (int32, error)
	Peek() (int32, error)
}

// Function to push elements to the stack
func (s *StackArray) Push(element int32) {
	s.Array = append(s.Array, 0)
	copy(s.Array[1:], s.Array)
	s.Array[0] = element
}

// Function to pop elements from the stack
// Returns error if stack is empty
func (s *StackArray) Pop() (int32, error) {

	if len(s.Array) == 0 {
		return 0, errors.New("Stack is empty!")
	}
	returnVal := s.Array[0]
	s.Array = s.Array[1:]
	return returnVal, nil

}

// Function to see the element at the top of the stack
// Returns error if stack is empty
func (s *StackArray) Peek() (int32, error) {
	if len(s.Array) == 0 {
		return 0, errors.New("Stack is empty!")
	}
	return s.Array[0], nil
}
