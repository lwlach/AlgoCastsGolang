package queue

import "fmt"

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

func (q *Queue) Peek() interface{} {
	if len(q.queue) == 0 {
		return nil
	}
	return q.queue[0]
}

func (q Queue) String() string {
	var output string
	for _, i := range q.queue {
		output += fmt.Sprintf("%v, ", i)
	}
	output = output[:len(output)-2]
	return output
}
