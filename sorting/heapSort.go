package sorting

func buildHeap(arr []int) {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
}

func heapify(arr []int, heapSize, i int) {
	l := 2 * i + 1
	r := 2 * i + 2
	largest := i
	if l < heapSize && arr[l] > arr[largest] {
		largest = l
	}
	if r < heapSize && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, heapSize, largest)
	}
}

func HeapSort(arr []int) {
	buildHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	} 
}