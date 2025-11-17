package day5

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	initialBuckets = 8
	loadFactor     = 1.0
	shrinkFactor   = 0.25
)

type bucket struct {
	key    string
	value  string
	active bool
}

type hashMap struct {
	buckets []bucket
	size    int
}

func NewHashMap() *hashMap {
	return &hashMap{
		buckets: make([]bucket, initialBuckets),
		size:    0,
	}
}

func (h *hashMap) hash(key string) int {
	hash := 0

	for _, ch := range key {
		hash += int(ch) % len(h.buckets)
	}

	return hash % len(h.buckets)
}

func (h *hashMap) load() float64 {
	return float64(h.size) / float64(len(h.buckets))
}

func (h *hashMap) resize(newCount int) {
	old := h.buckets
	h.buckets = make([]bucket, newCount)
	h.size = 0

	for _, b := range old {
		if b.active {
			h.insert(b.key, b.value)
		}
	}
}

func (h *hashMap) insert(key, value string) error {
	if h.load() >= loadFactor {
		h.resize(len(h.buckets) * 2)
	}

	idx := h.hash(key)

	for {
		if !h.buckets[idx].active {
			h.buckets[idx] = bucket{
				key:    key,
				value:  value,
				active: true,
			}
			h.size++
			return nil
		}

		if h.buckets[idx].key == key {
			h.buckets[idx].value = value
			return nil
		}

		idx = (idx + 1) % len(h.buckets)
	}
}

func (h *hashMap) get(key string) string {
	idx := h.hash(key)

	for {
		if !h.buckets[idx].active {
			return ""
		}

		if h.buckets[idx].key == key {
			return h.buckets[idx].value
		}

		idx = (idx + 1) % len(h.buckets)
	}
}

func (h *hashMap) delete(key string) error {
	idx := h.hash(key)

	for {
		if !h.buckets[idx].active {
			return errors.New("user not found")
		}

		if h.buckets[idx].key == key {
			h.buckets[idx].active = false
			h.size--

			if h.load() <= shrinkFactor && len(h.buckets) > initialBuckets {
				h.resize(len(h.buckets) / 2)
			}

			return nil
		}

		idx = (idx + 1) % len(h.buckets)
	}
}

func (h *hashMap) display() {
	if h.size == 0 {
		fmt.Println("Map is empty!")
		return
	}

	startBucket := rand.Intn(len(h.buckets))
	count := 0
	for b := 0; b < len(h.buckets); b++ {
		idx := (startBucket + b) % len(h.buckets)

		if h.buckets[idx].active {
			fmt.Printf("Key: %s, Value: %s\n",
				h.buckets[idx].key,
				h.buckets[idx].value,
			)
			count++
		}

		if count == h.size {
			break
		}
	}

	fmt.Println()
}

func (h *hashMap) authenticate(key, val string) bool {
	value := h.get(key)
	return val == value
}

func HashMap() {
	h := NewHashMap()

	h.insert("test", "123")
	h.insert("john", "abc")
	h.insert("mike", "pass")
	h.display()

	fmt.Printf("test: %s\n\n", h.get("test"))
	fmt.Printf("abc: %s\n\n", h.get("abc"))

	h.delete("john")

	h.display()

	fmt.Println("test User authentication:", h.authenticate("test", "123"))
	fmt.Println("john User authentication:", h.authenticate("john", "123"))
}
