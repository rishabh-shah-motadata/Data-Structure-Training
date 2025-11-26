package day9

import "fmt"

type wishItem struct {
	id   string
	name string
	next *wishItem
}

type wishlist struct {
	head *wishItem
}

func newWishlist() *wishlist {
	return &wishlist{}
}

func (w *wishlist) add(id, name string) {
	newItem := &wishItem{id, name, nil}

	if w.head == nil {
		w.head = newItem
		return
	}

	cur := w.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = newItem
}

func (w *wishlist) remove(id string) bool {
	if w.head == nil {
		return false
	}

	if w.head.id == id {
		w.head = w.head.next
		return true
	}

	prev := w.head
	cur := w.head.next

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

func (w *wishlist) display() {
	cur := w.head
	for cur != nil {
		fmt.Printf("ID: %s | Name: %s\n", cur.id, cur.name)
		cur = cur.next
	}
}
