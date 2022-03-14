package leetcode

type MyStack13 struct {
	queue []int
}

/** Initialize your data structure here. */
func ConstructorMyStack13() (s MyStack13) {
	return
}

/** Push element x onto stack. */
func (s *MyStack13) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack13) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

/** Get the top element. */
func (s *MyStack13) Top() int {
	return s.queue[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack13) Empty() bool {
	return len(s.queue) == 0
}
