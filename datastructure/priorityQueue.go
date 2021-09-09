package main

import (
	"fmt"
	"errors"
)

type PriorityQueue struct {
	arr []int
	heapSize int
}

func (pq *PriorityQueue) BuildHeap() {
	pq.heapSize = len(pq.arr)
	for i := pq.heapSize / 2 - 1; i >= 0; i-- {
		pq.Heapify(i)
	}
}

func (pq *PriorityQueue) Heapify(i int) {
	l := 2*i + 1
	r := 2*i + 2
	var largest = i
	if l < pq.heapSize && pq.arr[l] > pq.arr[largest] {
		largest = l
	}
	if r < pq.heapSize && pq.arr[r] > pq.arr[largest] {
		largest = r
	}
	if largest != i {
		pq.arr[i], pq.arr[largest] = pq.arr[largest], pq.arr[i]
		pq.Heapify(largest)
	}
}

func (pq *PriorityQueue) HeapMax() int {
	return pq.arr[0]
}

func (pq *PriorityQueue) ExtractMax() (int, error) {
	if pq.heapSize < 1 {
		return 0, errors.New("Queue is empty")
	}
	max := pq.arr[0]
	pq.arr[0] = pq.arr[pq.heapSize-1]
	pq.heapSize--
	pq.arr = pq.arr[:pq.heapSize]
	pq.Heapify(0)
	return max, nil
}

func (pq *PriorityQueue) HeapIncreaseKey(i, key int) error {
	if key < pq.arr[i] {
		return errors.New("Key is lower than current")
	}
	pq.arr[i] = key
	for i > 0 && pq.arr[(i-1)/2] < pq.arr[i] {
		parrent := (i-1)/2
		pq.arr[i], pq.arr[parrent] = pq.arr[parrent], pq.arr[i]
		i = parrent
	}
	return nil
}

func (pq *PriorityQueue) MaxHeapInsert(key int) {
	pq.arr = append(pq.arr, -100007)
	pq.heapSize++
	pq.HeapIncreaseKey(pq.heapSize-1, key)
}

func main() {
	pq := PriorityQueue{
		arr : []int{4,1,3,2,16,9,10,14,8,7},
	}
	fmt.Println(pq.arr)
	pq.BuildHeap()
	fmt.Println(pq.arr)
	max, _ := pq.ExtractMax()
	fmt.Println(max)
	fmt.Println(pq.arr)
	pq.HeapIncreaseKey(8,15)
	fmt.Println(pq.arr)
	pq.MaxHeapInsert(9)
	fmt.Println(pq.arr)
}