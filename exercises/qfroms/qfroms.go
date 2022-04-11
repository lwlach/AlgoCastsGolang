package qfroms

import "github.com/lwlach/AlgoCastsGolang/exercises/stack"

type Queue struct {
	s1, s2 stack.Stack
}

func (q *Queue) Add(i interface{}) {
	q.s1.Push(i)
}

func (q *Queue) Remove() interface{} {
	swap(&q.s1, &q.s2)
	result := q.s2.Pop()
	swap(&q.s2, &q.s1)
	return result
}

func (q *Queue) Peek() interface{} {
	swap(&q.s1, &q.s2)
	result := q.s2.Peek()
	swap(&q.s2, &q.s1)
	return result
}

func swap(from, to *stack.Stack) {
	for from.Peek() != nil {
		to.Push(from.Pop())
	}
}
