package main

import "fmt"

type Queue struct {
	items chan int
}

func (q *Queue) Enquqeue(i int) {
	q.items <- i
}

func (q *Queue) Dequeue() int {
	return <-q.items
}

func main() {
	q := Queue{
		items: make(chan int, 16),
	}

	q.Enquqeue(1)
	q.Enquqeue(2)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
