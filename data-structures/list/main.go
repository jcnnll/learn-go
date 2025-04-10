package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{
		value: value,
	}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.previous = l.tail
	}
	l.tail = node
}

type Node struct {
	value    int
	next     *Node
	previous *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Previous() *Node {
	return n.previous
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for {
		fmt.Println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}
	n = l.Last()
	for {
		fmt.Println(n.value)
		n = n.Previous()
		if n == nil {
			break
		}
	}
}
