package main

import (
	"fmt"
)

type SQ struct {
	// Реализуйте стек при помощи очереди
	queueLHS Queue
	queueRHS Queue
}

func (sq *SQ) push(item interface{}) {
	sq.queueLHS.enqueue(item)
}

func (sq *SQ) pop() interface{} {
	/*
	Перекладываю всё до последнего элемента во вторую очередь. Возвращаю единственный элемент из первой очереди.
	Меняю местами очереди
	*/
	if sq.queueLHS.top() == nil {
		return nil
	}
	var val interface{}
	for sq.queueLHS.top() != nil {
		val = sq.queueLHS.dequeue()
		if sq.queueLHS.top() == nil { break }
		sq.queueRHS.enqueue(val)		
	}
	sq.queueLHS, sq.queueRHS = sq.queueRHS, sq.queueLHS
	return val
}

type Queue struct {
	arr []interface{}
}

func (q *Queue) enqueue(item interface{}) {
	q.arr = append(q.arr, item)
}

func (q *Queue) dequeue() interface{} {
	var result interface{}
	if len(q.arr) > 0 {
		result = q.arr[0]
		q.arr = q.arr[1:]
		return result
	}
	return nil

}

func (q *Queue) top() interface{} {
	if len(q.arr) > 0 {
		return q.arr[0]
	}
	return nil
}

func (q *Queue) inverse(k int) Queue {
	// Обратите первые k элементов в очереди
	result := Queue {
		arr: make([]interface{}, 0),
	}
	sq := SQ{ // просто стек
		queueLHS : Queue{
			arr: make([]interface{}, 0),
		},
		queueRHS : Queue{
			arr: make([]interface{}, 0),
		},
	}
	/*
	к элементов суём в стек, потом достаём из него в новую очередь и 
	достаём оставшееся из оригинальной очереди
	*/
	for i := 0; i < k; i++ {
		sq.push(q.dequeue())
	}
	for i := 0; i < k; i++ {
		result.enqueue(sq.pop())
	}
	for q.top() != nil {
		result.enqueue(q.dequeue())
	}

	return result
}

func generateBin(n int) {
	// Сгенерируйте двоичные числа от 1 до n при помощи очереди
	q := Queue {
		arr : make([]interface{}, 0),
	}
	q.enqueue("1")
	for i := 0; i < n; i++ {
		val := q.top()
		fmt.Println(val)
		q.dequeue()
		q.enqueue(val.(string) + "0")
		q.enqueue(val.(string) + "1")
	}
}


func main() {
	q := Queue{
		arr: make([]interface{}, 0),
	}
	
	q.enqueue(3)
	q.enqueue(5)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	
	sq := SQ{
		queueLHS : Queue{
			arr: make([]interface{}, 0),
		},
		queueRHS : Queue{
			arr: make([]interface{}, 0),
		},
	}
	fmt.Println(sq.pop())
	sq.push(4)
	sq.push(6)
	sq.push(8)
	fmt.Println(sq.pop())
	fmt.Println(sq.pop())
	fmt.Println(sq.pop())

	fmt.Println("----------")

	q.enqueue(3)
	q.enqueue(5)
	q.enqueue(2)
	q.enqueue(1)
	q = q.inverse(4)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())

	fmt.Println("----------")

	generateBin(10)
}
