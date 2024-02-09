package collections

import (
	"testing"
)

func (ta TestAssigner) AssignPriority(item interface{}) func() float32 {
	return func() float32 {
		return ta.weight * 1
	}
}

func TestOrderedList_Insert(t *testing.T) {
	// Create a new OrderedList
	ol := NewOrderedList(TestAssigner{weight: 1.0})

	// Insert some values with priorities
	ol.Insert("value1")
	ol.Insert("value2")
	ol.Insert("value3")

	// Verify the order of the values
	expectedOrder := []string{"value3", "value1", "value2"}
	for i, val := range expectedOrder {
		if ol.Get(i) != val {
			t.Errorf("Expected value at index %d to be %s, got %s", i, val, ol.Get(i))
		}
	}
}

func TestOrderedList_Get(t *testing.T) {
	// Create a new OrderedList
	ol := NewOrderedList()

	// Insert some values with priorities
	ol.Insert("value1", 1.0)
	ol.Insert("value2", 2.0)
	ol.Insert("value3", 0.5)

	// Verify the values retrieved by index
	expectedValues := []string{"value3", "value1", "value2"}
	for i, val := range expectedValues {
		if ol.Get(i) != val {
			t.Errorf("Expected value at index %d to be %s, got %s", i, val, ol.Get(i))
		}
	}
}

// Add more test cases as needed
