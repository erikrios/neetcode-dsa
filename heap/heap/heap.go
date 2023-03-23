package heap

type Heap struct {
	values []int
}

func New() *Heap {
	return &Heap{values: []int{0}}
}

func (h *Heap) Push(v int) {
	h.values = append(h.values, v)
	i := len(h.values) - 1

	for h.values[i] < h.values[i/2] {
		h.values[i], h.values[i/2] = h.values[i/2], h.values[i]
		i /= 2
	}
}

func (h *Heap) Pop() int {
	if len(h.values) == 1 {
		return 0
	}

	if len(h.values) == 2 {
		return h.Pop()
	}

	res := h.values[1]
	// Move last value to root
	h.values[1] = h.Pop()
	i := 1

	for 2*i < len(h.values) {
		if 2*i+1 < len(h.values) && h.values[2*i+1] < h.values[2*i] && h.values[i] > h.values[2*i+1] {
			// Swap right child
			h.values[i], h.values[2*i+1] = h.values[2*i+1], h.values[i]
			i = 2*i + 1
		} else if h.values[i] > h.values[2*i] {
			// Swap left child
			h.values[i], h.values[2*i] = h.values[2*i], h.values[i]
			i = 2 * i
		} else {
			break
		}
	}

	return res
}

func (h *Heap) Heapify(nums []int) {
	// 0-th position is moved to the end
	nums = append(nums, nums[0])

	h.values = nums
	cur := len(h.values) - 1/2

	for cur > 0 {
		i := cur
		if 2*i+1 < len(h.values) && h.values[2*i+1] < h.values[2*i] && h.values[i] > h.values[2*i+1] {
			// Swap right child
			h.values[i], h.values[2*i+1] = h.values[2*i+1], h.values[i]
			i = 2*i + 1
		} else if h.values[i] > h.values[2*i] {
			// Swap left child
			h.values[i], h.values[2*i] = h.values[2*i], h.values[i]
			i = 2 * i
		} else {
			break
		}
	}

	cur--
}
