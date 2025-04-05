package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	count := len(s.items)

	if count == 0 {
		return -1
	}
	item, items := s.items[count-1], s.items[0:count-1]
	s.items = items
	return item
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
