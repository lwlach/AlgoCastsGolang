package weave

import "github.com/lwlach/AlgoCastsGolang/exercises/queue"

func Weave(q1, q2 queue.Queue) queue.Queue {
	var q = queue.Queue{}
	for q1.Peek() != nil || q2.Peek() != nil {
		if q1.Peek() != nil {
			q.Add(q1.Remove())
		}
		if q2.Peek() != nil {
			q.Add(q2.Remove())
		}
	}

	return q
}
