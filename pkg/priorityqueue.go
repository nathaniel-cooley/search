package search

import "container/heap"

// Item represents an item in the priority queue.
type Item struct {
	value       interface{}
	getPriority func() int
	index       int
}

type Assigner interface {
	AssignPriority(item Item) func() int
}

// PriorityQueue implements a priority queue.
type PriorityQueue struct {
	items    []Item
	assigner Assigner
}

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int {
	return len(pq.items)
}

// Less checks if the item at index i has higher priority than the item at index j.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.items[i].getPriority() < pq.items[j].getPriority()
}

// Swap swaps the items at indexes i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = j
	pq.items[j].index = i
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := Item{value: x, index: n}
	item.getPriority = pq.assigner.AssignPriority(item)
	pq.items = append(pq.items, item)
	heap.Fix(pq, item.index)
}

// Pop removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[0]
	item.index = -1 // for safety

	pq.Swap(0, n-1)
	pq.items = old[0 : n-1]
	heap.Fix(pq, 0)
	return item
}

// Update modifies the priority and value of an item in the priority queue.
func (pq *PriorityQueue) Update(item *Item, value interface{}) {
	item.value = value
	pq.assigner.AssignPriority(*item)
	heap.Fix(pq, item.index)
}
