package day2

import (
	"fmt"
)

const (
	initialSize = 2
)

type dynamicArray struct {
	data []string
	len  int
	cap  int
}

func newDynamicArray() *dynamicArray {
	return &dynamicArray{
		data: (make([]string, initialSize)),
		len:  0,
		cap:  initialSize,
	}
}

func (d *dynamicArray) add(val string) {
	if d.len == d.cap {
		d.resize()
	}

	d.data[d.len] = val
	d.len++
}

func (d *dynamicArray) resize() {
	newCap := d.cap * 2
	newData := make([]string, newCap)

	for i := 0; i < d.len; i++ {
		newData[i] = d.data[i]
	}

	d.data = newData
	d.cap = newCap
}

func (d *dynamicArray) print() {
	fmt.Printf("Dynamic Array is:\nlen=%d\ncap=%d\ndArr=> ", d.len, d.cap)
	for i := 0; i < d.len; i++ {
		fmt.Printf("%s ", d.data[i])
	}
	fmt.Printf("\n\n")
}

func (d *dynamicArray) remove() {
	if d.len == 0 {
		fmt.Println("Array is empty, nothing to remove.")
		return
	}

	d.len--
	d.data[d.len] = ""

	if d.len > 0 && d.len <= d.cap/4 {
		d.shrink()
	}
}

func (d *dynamicArray) shrink() {
	newCap := max(d.cap/2, initialSize)

	newData := make([]string, newCap)
	for i := 0; i < d.len; i++ {
		newData[i] = d.data[i]
	}

	d.data = newData
	d.cap = newCap
}

func DynamicArrayImplementation() {
	dArr := newDynamicArray()

	packets := [6]string{
		"packet-1", "packet-2", "packet-3", "packet-4", "packet-5", "packet-6",
	}

	for i := range packets {
		dArr.add(packets[i])
		dArr.print()
	}

	dArr.remove()
	dArr.print()

	dArr.remove()
	dArr.print()

	dArr.remove()
	dArr.print()

	dArr.remove()
	dArr.print()
}
