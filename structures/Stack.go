package structures

// Stack of int
type Stack struct {
	nums []int
}

// NewStack retrun *Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push n to Stack
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop from stack at top
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// Len of now stack
func (s *Stack) Len() int {
	return len(s.nums)
}

// IsEmpty now stack is empty?
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
