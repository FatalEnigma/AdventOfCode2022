package types

type Stack []interface{}

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, element) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// Reverse reverse the order of the stack
func (s *Stack) Reverse() {
	temp := Stack{}

	for !s.IsEmpty() {
		value, _ := s.Pop()
		temp.Push(value)
	}

	*s = temp
}

// Transfer moves num items to a destination stack
func (s *Stack) Transfer(destination *Stack, num int) int {
	var numMoved int

	for i := 0; i < num && !s.IsEmpty(); i++ {
		value, _ := s.Pop()
		destination.Push(value)
		numMoved++
	}

	return numMoved
}

// Copy returns a copy of the stack
func (s *Stack) Copy() Stack {
	stackCopy := *s
	temp := Stack{}

	for !stackCopy.IsEmpty() {
		value, _ := stackCopy.Pop()
		temp.Push(value)
	}

	temp.Reverse()

	return temp
}
