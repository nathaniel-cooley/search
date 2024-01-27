package search

import "container/heap"

// Item represents an item in the priority queue.
type ItemMaker interface {
	NewItem(value interface{}, priorityAssigner func() int, index int) Item
}

type Item struct {
	value       interface{}
	GetPriority func() int
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
	return pq.items[i].GetPriority() < pq.items[j].GetPriority()
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
	item := x.(Item)
	item.index = n
	pq.items = append(pq.items, item)
}

// Pop removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	pq.items = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an item in the priority queue.
func (pq *PriorityQueue) Update(item *Item, value interface{}) {
	item.value = value
	pq.assigner.AssignPriority(*item)
	heap.Fix(pq, item.index)
}
