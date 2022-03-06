package stack

type Stack struct {
	stack []interface{}
}

func (s *Stack) Push(i interface{}) {
	s.stack = append(s.stack, i)
}

func (s *Stack) Pop() interface{} {
	stackLen := len(s.stack)
	if len(s.stack) == 0 {
		return nil
	}
	top := s.stack[stackLen-1]
	s.stack[stackLen-1] = nil
	s.stack = s.stack[:stackLen-1]
	return top
}

func (s Stack) Peek() interface{} {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}
