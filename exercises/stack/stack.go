package stack

type Stack struct {
	stack []interface{}
}

func (s *Stack) Push(i interface{}) {
	s.stack = append(s.stack, i)
}

func (s *Stack) Pop() interface{} {
	if len(s.stack) == 0 {
		return nil
	}
	last := len(s.stack) - 1
	top := s.stack[last]
	s.stack[last] = nil
	s.stack = s.stack[:last]
	return top
}

func (s Stack) Peek() interface{} {
	if len(s.stack) == 0 {
		return nil
	}
	last := len(s.stack) - 1
	return s.stack[last]
}
