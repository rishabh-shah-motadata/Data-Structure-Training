package day9

import "fmt"

type discounts struct {
	items map[string]float64
}

func newDiscount() *discounts {
	return &discounts{items: make(map[string]float64)}
}

func (d *discounts) add(id string, discount float64) {
	d.items[id] = discount
}

func (d *discounts) remove(id string) {
	delete(d.items, id)
}

func (d *discounts) display() {
	for id, discount := range d.items {
		fmt.Printf("ID: %s | Discount: %0.2f\n", id, discount)
	}
}
