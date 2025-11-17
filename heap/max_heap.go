package heap

import (
	"fmt"
)

type maxHeap struct {
	data []int
}

func (h *maxHeap) insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *maxHeap) extractMax() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	max := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	h.heapifyDown(0)

	return max, true
}

func (h *maxHeap) peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *maxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] >= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

func (h *maxHeap) heapifyDown(index int) {
	n := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

func MaxHeap() {
	h := &maxHeap{}

	nums := []int{3, 2, 1, 5, 4}
	for _, n := range nums {
		fmt.Printf("Inserting %d → ", n)
		h.insert(n)
		fmt.Println(h.data)
	}

	val, _ := h.peek()
	fmt.Printf("\nPeek (Max Element): %d\n", val)

	fmt.Println("\nExtracting elements (descending order):")
	for len(h.data) > 0 {
		max, _ := h.extractMax()
		fmt.Printf("Extracted %d → Heap: %v\n", max, h.data)
	}

	fmt.Printf("\n\n")
}
