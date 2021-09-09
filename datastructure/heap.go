package main

import (
	"fmt"
)

func BuildMaxHeap(arr []int) {
	heapSize := len(arr)
	for i := heapSize / 2 - 1; i >= 0; i-- {
		MaxHeapify(arr, heapSize, i)
	}
}

func MaxHeapify(arr []int, heapSize, i int) {
	l := 2*i + 1
	r := 2*i + 2
	var largest int
	if l < heapSize && arr[l] > arr[i] { // change on - for minHeap
		largest = l
	} else {
		largest = i
	}
	if r < heapSize && arr[r] > arr[largest] { // change on - for minHeap
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		MaxHeapify(arr, heapSize, largest)
	}
}


func HeapSort(arr []int) {
	BuildMaxHeap(arr)
	heapSize := len(arr)
	for i := heapSize - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		MaxHeapify(arr, i, 0)
	}
}

func main() {
	nums := []int{4,1,3,2,16,9,10,14,8,7, 0}
	// BuildMaxHeap(nums)
	HeapSort(nums)
	fmt.Println(nums)
}
