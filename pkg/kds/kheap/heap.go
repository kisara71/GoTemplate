package kheap

import "github.com/kisara71/GoTemplate/pkg/kbase"

// Comparator is a function type that defines the comparison logic for two elements.
// It should return true if 'a' is considered "better" than 'b' based on the heap's ordering.
// For a min-heap, it would be 'a < b'. For a max-heap, it would be 'a > b'.
type Comparator[T any] func(a, b T) bool

// KHeap is a generic heap data structure.
// It maintains the heap property based on the provided Comparator.
type KHeap[T any] struct {
	Data []T
	CMP  Comparator[T]
}

// NewKHeap creates and initializes a new KHeap from a given slice of data.
// It builds the heap by heapifying down from the middle of the slice.
func NewKHeap[T any](data []T, cmp Comparator[T]) *KHeap[T] {
	heap := &KHeap[T]{
		Data: data,
		CMP:  cmp,
	}

	// Build heap: start from the last non-leaf node and heapify down.
	for i := len(data)/2 - 1; i >= 0; i-- {
		heap.heapifyDown(i)
	}
	return heap
}

// heapifyDown maintains the heap property by moving an element down the heap.
// It compares the element at 'idx' with its children and swaps it with the "best" child
// (according to the comparator) if the heap property is violated.
func (h *KHeap[T]) heapifyDown(idx int) {
	best := idx       // Assume current node is the best
	lson := idx*2 + 1 // Left child index
	rson := idx*2 + 2 // Right child index

	// Check if left child exists and is "better" than the current best
	if lson < len(h.Data) && h.CMP(h.Data[lson], h.Data[best]) {
		best = lson
	}
	// Check if right child exists and is "better" than the current best
	if rson < len(h.Data) && h.CMP(h.Data[rson], h.Data[best]) {
		best = rson
	}
	// If a child is better, swap and recursively heapify down
	if best != idx {
		kbase.Swap(&h.Data[best], &h.Data[idx])
		h.heapifyDown(best)
	}
}

// heapifyUp maintains the heap property by moving an element up the heap.
// It compares the element at 'idx' with its parent and swaps them
// if the heap property is violated, continuing until the root or property holds.
func (h *KHeap[T]) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2 // Calculate parent index
		// If current node is "better" than its parent, swap them
		if h.CMP(h.Data[idx], h.Data[parent]) {
			kbase.Swap(&h.Data[parent], &h.Data[idx])
			idx = parent // Move up to the parent's position
		} else {
			break // Heap property holds, stop
		}
	}
}

// Push adds a new value to the heap and maintains the heap property.
func (h *KHeap[T]) Push(val T) {
	h.Data = append(h.Data, val) // Add new element to the end
	h.heapifyUp(h.Size() - 1)    // Heapify up from the new element's position
}

// Top returns the top element of the heap without removing it.
// Returns the element and true if the heap is not empty, otherwise returns a zero value and false.
func (h *KHeap[T]) Top() (T, bool) {
	if !h.Empty() {
		return h.Data[0], true // Top element is always at index 0
	}
	var zero T // Zero value for type T
	return zero, false
}

// Empty checks if the heap is empty.
func (h *KHeap[T]) Empty() bool {
	return len(h.Data) == 0
}

// Pop removes and returns the top element from the heap.
// Returns the element and true if the heap is not empty, otherwise returns a zero value and false.
func (h *KHeap[T]) Pop() (T, bool) {
	var val T
	if h.Empty() {
		return val, false
	}
	val = h.Data[0]                             // Get the top element
	kbase.Swap(&h.Data[0], &h.Data[h.Size()-1]) // Swap top with last element
	h.Data = h.Data[:h.Size()-1]                // Remove the last element (original top)
	h.heapifyDown(0)                            // Heapify down from the new root
	return val, true
}

// Size returns the number of elements in the heap.
func (h *KHeap[T]) Size() int {
	return len(h.Data)
}

// Clear removes all elements from the heap.
func (h *KHeap[T]) Clear() {
	h.Data = h.Data[:0]
}
