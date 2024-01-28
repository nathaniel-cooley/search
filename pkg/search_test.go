package search

import (
	"testing"
)

type TestAssigner struct {
	weight int
}

func (ta TestAssigner) AssignPriority(item Item) func() int {
	return func() int {
		return ta.weight * item.index
	}
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: 1}}

	pq.Push("item1")
	pq.Push("item2")
	pq.Push("item3")

	if pq.Len() != 3 {
		t.Errorf("PriorityQueue.Len() = %v, want %v", pq.Len(), 3)
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: -1}}
	pq.Push("item1")
	pq.Push("item2")

	if pq.Less(1, 0) {
		t.Errorf("PriorityQueue.Less() = %v, want %v", pq.Less(0, 1), false)
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	// TODO: Write test cases for Swap method
}

func TestPriorityQueue_Push(t *testing.T) {
	// TODO: Write test cases for Push method
}

func TestPriorityQueue_Pop(t *testing.T) {
	// TODO: Write test cases for Pop method
}

func TestPriorityQueue_Update(t *testing.T) {
	// TODO: Write test cases for Update method
}
