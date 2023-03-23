package hashmap

import "reflect"

type pair struct {
	key string
	val string
}

type HashMap struct {
	size     int
	capacity int
	maps     []pair
}

func New() *HashMap {
	return &HashMap{size: 0, capacity: 2, maps: []pair{{}, {}}}
}

func (h *HashMap) Hash(key string) int {
	var index int
	for i := 0; i < len(key); i++ {
		index += int(key[i])
	}

	return index % h.capacity
}

func (h *HashMap) Get(key string) string {
	index := h.Hash(key)

	for !reflect.DeepEqual(h.maps[index], pair{}) {
		if h.maps[index].key == key {
			return h.maps[index].val
		}
		index++
		index %= h.capacity
	}

	return ""
}

func (h *HashMap) Put(key, val string) {
	index := h.Hash(key)

	for {
		if reflect.DeepEqual(h.maps[index], pair{}) {
			h.maps[index] = pair{key: key, val: val}
			h.size++
			if h.size >= h.capacity/2 {
				h.rehash()
			}
			return
		} else if h.maps[index].key == key {
			h.maps[index].val = val
			return
		}

		index++
		index %= h.capacity
	}
}

func (h *HashMap) rehash() {
	h.capacity *= 2
	newMaps := make([]pair, h.capacity, h.capacity)

	oldMaps := h.maps
	h.maps = newMaps
	h.size = 0
	for i := 0; i < len(oldMaps); i++ {
		pairV := oldMaps[i]
		if !reflect.DeepEqual(pairV, pair{}) {
			h.Put(pairV.key, pairV.val)
		}
	}
}
