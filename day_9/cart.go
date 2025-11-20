package day9

import "fmt"

type item struct {
	id       string
	name     string
	price    float64
	quantity int
	next     *item
}

type cart struct {
	head *item
}

func (c *cart) add(id, name string, price float64, qty int) {
	newItem := &item{id, name, price, qty, nil}

	if c.head == nil {
		c.head = newItem
		return
	}

	cur := c.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newItem
}

func (c *cart) remove(id string) bool {
	if c.head == nil {
		return false
	}

	if c.head.id == id {
		c.head = c.head.next
		return true
	}

	prev := c.head
	cur := c.head.next

	for cur != nil {
		if cur.id == id {
			prev.next = cur.next
			return true
		}
		prev = cur
		cur = cur.next
	}

	return false
}

func (c *cart) display() {
	cur := c.head
	for cur != nil {
		fmt.Printf("ID: %s | Name: %s | Price: %.2f | Qty: %d\n",
			cur.id, cur.name, cur.price, cur.quantity)
		cur = cur.next
	}
}
