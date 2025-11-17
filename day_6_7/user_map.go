package day_6_7

type user struct {
	username string
	password string
}

type bucket struct {
	key    string
	value  user
	active bool
}

type userHashMap struct {
	buckets []bucket
	size    int
}

func newUserHashMap() *userHashMap {
	return &userHashMap{
		buckets: make([]bucket, 8),
		size:    0,
	}
}

func (h *userHashMap) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (hash*31 + int(ch)) % len(h.buckets)
	}
	return hash
}

func (h *userHashMap) Put(key string, value user) {
	idx := h.hash(key)

	for {
		if !h.buckets[idx].active {
			h.buckets[idx] = bucket{key: key, value: value, active: true}
			h.size++
			return
		} else if h.buckets[idx].key == key {
			h.buckets[idx].value = value
			return
		}
		idx = (idx + 1) % len(h.buckets)
	}
}

func (h *userHashMap) Get(key string) (user, bool) {
	idx := h.hash(key)

	for {
		if !h.buckets[idx].active {
			return user{}, false
		}
		if h.buckets[idx].key == key {
			return h.buckets[idx].value, true
		}
		idx = (idx + 1) % len(h.buckets)
	}
}
