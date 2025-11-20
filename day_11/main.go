package day11

import "fmt"

type task struct {
	name     string
	priority int
}

type taskMaxHeap struct {
	data []task
}

func newMaxHeap(capacity int) *taskMaxHeap {
    return &taskMaxHeap{
        data: make([]task, 0, capacity),
    }
}

func (h *taskMaxHeap) insert(t task) {
	h.data = append(h.data, t)
	h.heapifyUp(len(h.data) - 1)
}

func (h *taskMaxHeap) extractMax() (task, bool) {
	if len(h.data) == 0 {
		return task{}, false
	}

	max := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	h.heapifyDown(0)

	return max, true
}

func (h *taskMaxHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent].priority >= h.data[i].priority {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *taskMaxHeap) heapifyDown(i int) {
	n := len(h.data)

	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < n && h.data[left].priority > h.data[largest].priority {
			largest = left
		}
		if right < n && h.data[right].priority > h.data[largest].priority {
			largest = right
		}

		if largest == i {
			break
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

func TaskSchedulerDemo() {
	h := newMaxHeap(20)

	tasks := []task{
		{"Update Inventory", 2},
		{"Send Email", 1},
		{"Generate Reports", 4},
		{"Process Payment", 5},
		{"Backup Database", 3},
	}

	fmt.Println("\nInserting tasks:")
	for _, t := range tasks {
		fmt.Printf("Adding %-18s (Priority %d)\n", t.name, t.priority)
		h.insert(t)
	}

	fmt.Println("\nExecuting tasks by priority:")
	for len(h.data) > 0 {
		t, _ := h.extractMax()
		fmt.Printf("Executing: %-18s | Priority: %d\n", t.name, t.priority)
	}
}
