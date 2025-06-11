package ps

// Stack represents a simple stack with the top at index 0.
type Stack struct {
	data []int
}

// NewStack creates a stack from a slice (top element first).
func NewStack(nums []int) *Stack {
	d := make([]int, len(nums))
	copy(d, nums)
	return &Stack{data: d}
}

func (s *Stack) Len() int { return len(s.data) }

// Push adds an element on the top of the stack.
func (s *Stack) Push(v int) {
	s.data = append([]int{v}, s.data...)
}

// Pop removes the top element of the stack and returns it.
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	v := s.data[0]
	s.data = s.data[1:]
	return v, true
}

// Swap swaps the first two elements of the stack.
func (s *Stack) Swap() bool {
	if len(s.data) < 2 {
		return false
	}
	s.data[0], s.data[1] = s.data[1], s.data[0]
	return true
}

// Rotate shifts up all elements by 1: the first becomes last.
func (s *Stack) Rotate() bool {
	if len(s.data) < 1 {
		return false
	}
	first := s.data[0]
	s.data = append(s.data[1:], first)
	return true
}

// ReverseRotate shifts down all elements by 1: the last becomes first.
func (s *Stack) ReverseRotate() bool {
	if len(s.data) < 1 {
		return false
	}
	last := s.data[len(s.data)-1]
	s.data = append([]int{last}, s.data[:len(s.data)-1]...)
	return true
}

// Peek returns top element without removing.
func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[0], true
}

// Data returns the underlying slice (copy).
func (s *Stack) Data() []int {
	d := make([]int, len(s.data))
	copy(d, s.data)
	return d
}
