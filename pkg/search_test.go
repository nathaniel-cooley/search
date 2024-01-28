package search

import (
	"testing"
)

type TestAssigner struct {
	weight int
}

func (ta TestAssigner) AssignPriority(item Item) func() int {
	return func() int {
		return ta.weight * item.value.(int)
	}
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: 1}}

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	if pq.Len() != 3 {
		t.Errorf("PriorityQueue.Len() = %v, want %v", pq.Len(), 3)
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: -1}}
	expectedValue := 1
	pq.Push(2)
	pq.Push(expectedValue)

	// index 0 has higher priority than index 1, so index 0 should be less in a min heap
	if pq.Less(1, 0) && pq.items[0].value.(int) == expectedValue {
		t.Errorf("PriorityQueue.Less() = %v, want %v", pq.Less(1, 0), false)
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: 1}}
	pq.Push(1)
	pq.Push(2)
	// before the values are swapped, the first item should be 1 and the second item should be 2

	pq.Swap(0, 1)

	if pq.items[0].value.(int) != 2 {
		t.Errorf("PriorityQueue.Swap() = %v, want %v", pq.items[0].value.(int), 2)
	}
	if pq.items[1].value.(int) != 1 {
		t.Errorf("PriorityQueue.Swap() = %v, want %v", pq.items[1].value.(int), 1)
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: 1}}
	pq.Push(3)
	pq.Push(2)
	pq.Push(1)

	if pq.items[0].value.(int) != 1 {
		t.Errorf("PriorityQueue.Push() = %v, want %v", pq.items[0].value.(int), 1)
	}
	if pq.items[1].value.(int) != 3 {
		t.Errorf("PriorityQueue.Push() = %v, want %v", pq.items[1].value.(int), 3)
	}
	if pq.items[2].value.(int) != 2 {
		t.Errorf("PriorityQueue.Push() = %v, want %v", pq.items[2].value.(int), 1)
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: 1}}
	pq.Push(3)
	pq.Push(4)
	pq.Push(2)
	pq.Push(1)

	for i := 0; i < 3; i++ {
		item := pq.Pop()
		if item.(Item).value.(int) != i+1 {
			t.Errorf("PriorityQueue.Pop() = %v, want %v", item.(Item).value.(int), i+1)
		}
	}
}

func TestPriorityQueue_Update(t *testing.T) {
	pq := PriorityQueue{assigner: TestAssigner{weight: 1}}
	pq.Push(3)
	pq.Push(4)
	pq.Push(2)
	pq.Push(1)

	pq.Update(&pq.items[0], 5)

	if pq.items[0].value.(int) != 5 {
		t.Errorf("PriorityQueue.Update() = %v, want %v", pq.items[0].value.(int), 5)
	}
}
