package day_6_7

import "fmt"

type messageNode struct {
	text string
	next *messageNode
}

type messageHistory struct {
	head *messageNode
}

func (m *messageHistory) addMessage(msg string) {
	newNode := &messageNode{text: msg}
	if m.head == nil {
		m.head = newNode
		return
	}
	
	temp := m.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (m *messageHistory) print() {
	temp := m.head
	for temp != nil {
		fmt.Println("•", temp.text)
		temp = temp.next
	}
}
