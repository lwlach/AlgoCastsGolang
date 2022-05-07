package main

import (
	"fmt"

	"github.com/lwlach/AlgoCastsGolang/exercises/queue"
	"github.com/lwlach/AlgoCastsGolang/exercises/weave"
)

func main() {
	q := queue.Queue{}
	q.Add("test")
	q.Add("t1")
	q.Add("t2")
	q.Add("t3")

	intQueue := queue.Queue{}
	intQueue.Add(1)
	intQueue.Add(2)
	intQueue.Add(3)

	q3 := weave.Weave(q, intQueue)
	fmt.Println(q3)
}
