package structures

// Queue save int list
type Queue struct {
	nums []int
}

// NewQueue returns a new *Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Push n to Queue
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop from Queue
func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

// Len of Queue
func (q *Queue) Len() int {
	return len(q.nums)
}

// IsEmpty returns Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
