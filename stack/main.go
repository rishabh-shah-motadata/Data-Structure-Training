package stack

import "fmt"

type node struct {
	data int
	next *node
}

type stack struct {
	top *node
	len int
}

func (s *stack) push(value int) {
	newNode := &node{data: value, next: s.top}
	s.top = newNode
	s.len++
}

func (s *stack) pop() {
	if s.top == nil {
		fmt.Println("Stack is empty!")
		return
	}

	value := s.top.data
	s.top = s.top.next
	s.len--
	fmt.Printf("Popped %d from stack\n", value)
}

func (s *stack) peek() {
	if s.top == nil {
		fmt.Println("Stack is empty!")
		return
	}

	fmt.Println("The top element is: ", s.top.data)
}

func (s *stack) display() {
	if s.top == nil {
		fmt.Println("Stack is empty!")
		return
	}

	fmt.Print("Stack elements are: ")
	for current := s.top; current != nil; current = current.next {
		fmt.Printf("%d ", current.data)
	}
	fmt.Println()
}

func Stack() {
	stack := &stack{}

	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.display()

	stack.peek()

	stack.pop()
	stack.display()

	fmt.Printf("Stack size: %d\n\n", stack.len)
}
