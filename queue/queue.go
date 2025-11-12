package queue

import "fmt"

type node struct {
	data int
	next *node
}

type queue struct {
	front *node
	rear  *node
	len   int
}

func (q *queue) enqueue(value int) {
	newNode := &node{data: value}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.len++
}

func (q *queue) dequeue() {
	if q.front == nil {
		fmt.Println("Queue is empty!")
		return
	}

	value := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	q.len--
	fmt.Printf("Dequeued %d from queue\n", value)
}

func (q *queue) display() {
	if q.front == nil {
		fmt.Println("Queue is empty!")
		return
	}

	fmt.Print("Queue elements are: ")
	for current := q.front; current != nil; current = current.next {
		fmt.Printf("%d ", current.data)
	}
	fmt.Println()
}

func Queue() {
	queue := &queue{}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.display()

	queue.dequeue()
	queue.display()

	fmt.Printf("Queue size: %d\n\n", queue.len)
}
