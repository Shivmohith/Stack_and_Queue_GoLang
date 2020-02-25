package queue

import "errors"

// A structure which has a queue array
type QueueArray struct {
	Array []int32
}

// Function to create a new queue
func NewQueueArray() *QueueArray {
	return &QueueArray{Array: make([]int32, 0)}
}

type QueueMethods interface {
	Push(int32)
	Pop() (int32, error)
	Peek() (int32, error)
}

// Function to insert an element into the queue
func (q *QueueArray) Push(element int32) {
	q.Array = append(q.Array, element)
}

// Function to remove an element from the queue
// Returns an error if the queue is empty
func (q *QueueArray) Pop() (int32, error) {
	if len(q.Array) == 0 {
		return 0, errors.New("Queue is empty!")
	}
	returnValue := q.Array[0]
	q.Array = q.Array[1:]
	return returnValue, nil

}

// Function to view the first element in the queue
// Returns an error if the queue is empty
func (q *QueueArray) Peek() (int32, error) {
	if len(q.Array) == 0 {
		return 0, errors.New("Queue is emplty!")
	}
	return q.Array[0], nil
}
