package queue

type Queue struct {
	queue []interface{}
}

func (q *Queue) Add(i interface{}) {
	q.queue = append(q.queue, i)
}

func (q *Queue) Remove() interface{} {
	if len(q.queue) == 0 {
		return nil
	}
	r := q.queue[0]
	q.queue[0] = nil
	q.queue = q.queue[1:]
	return r
}
