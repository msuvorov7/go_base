package main

import (
	"fmt"
)

type List struct {
	length int
	head *Node
	tail *Node
}

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

func (l *List) pushFront(item interface{}) {
	node := &Node {
		data: item,
		prev: nil,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.length++
}

func (l *List) pushBack(item interface{}) {
	node := &Node {
		data: item,
		prev: nil,
		next: nil,
	}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

func (l *List) print() {
	node := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	l := List {
		length: 0,
		head: nil,
		tail: nil,
	}
	l.pushBack(4)
	l.pushBack(5)
	l.pushFront(3)
	l.pushFront(2)
	l.print()
}