package weave

func Weave(q1, q2 Queue) Queue {
	var q = Queue{}
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
