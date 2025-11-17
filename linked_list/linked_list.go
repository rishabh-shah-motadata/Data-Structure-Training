package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	len  int
}

func (ll *linkedList) insert(value, index int) {
	if index < 0 || index > ll.len {
		fmt.Println("Invalid position!")
		return
	}

	newNode := &node{data: value}

	// Insert at the beginning
	if index == 0 {
		newNode.next = ll.head
		ll.head = newNode
		ll.len++
		return
	}

	current := ll.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
	ll.len++
}

func (ll *linkedList) delete(position int) {
	if position < 0 || position >= ll.len {
		fmt.Println("Invalid position!")
		return
	}

	// Delete from beginning
	if position == 0 {
		fmt.Printf("Deleted %d from position %d\n", ll.head.data, position)
		ll.head = ll.head.next
		ll.len--
		return
	}

	current := ll.head
	for i := 0; i < position-1; i++ {
		current = current.next
	}

	fmt.Printf("Deleted %d from position %d\n", current.next.data, position)
	current.next = current.next.next
	ll.len--
}

func (ll *linkedList) display() {
	if ll.head == nil {
		fmt.Println("Linked List is empty!")
		return
	}

	fmt.Print("Linked List elements are: ")
	for current := ll.head; current != nil; current = current.next {
		fmt.Printf("%d ", current.data)
	}
	fmt.Println()
}

// Demo usage
func LinkedList() {
	ll := &linkedList{}

	ll.insert(10, 0)
	ll.insert(20, 1)
	ll.insert(30, 2)
	ll.insert(5, 0)
	ll.insert(15, 3)
	ll.display()

	ll.delete(3)
	ll.display()

	fmt.Printf("Linked List size: %d\n\n", ll.len)
}
