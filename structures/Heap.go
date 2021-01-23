package structures

// intHeap for heap interface
type intHeap []int

// Push *h, to append h
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop *h, to decrease h
func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// Len of this heap
func (h intHeap) Len() int {
	return len(h)
}

// Less of i, j at this heap, return bool
func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap of i, j at this heap
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
