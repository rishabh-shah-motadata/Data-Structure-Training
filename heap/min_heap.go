package heap

import (
	"fmt"
)

type minHeap struct {
	data []int
}

func (h *minHeap) insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *minHeap) extractMin() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	min := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	h.heapifyDown(0)

	return min, true
}

func (h *minHeap) peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *minHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] <= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

func (h *minHeap) heapifyDown(index int) {
	n := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

func MinHeap() {
	h := &minHeap{}

	nums := []int{3, 2, 1, 5, 4}
	for _, n := range nums {
		fmt.Printf("Inserting %d → ", n)
		h.insert(n)
		fmt.Println(h.data)
	}

	val, _ := h.peek()
	fmt.Printf("\nPeek (Min Element): %d\n", val)

	fmt.Println("\nExtracting elements (ascending order):")
	for len(h.data) > 0 {
		min, _ := h.extractMin()
		fmt.Printf("Extracted %d → Heap: %v\n", min, h.data)
	}
	
	fmt.Printf("\n\n")
}
