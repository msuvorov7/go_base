package sorting

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lhs, rhs int) {
	if lhs > rhs {
		return
	}
	p := partition(arr, lhs, rhs)
	quickSort(arr, lhs, p-1)
	quickSort(arr, p+1, rhs)
}

func partition(arr []int, lhs, rhs int) int {
	x := arr[rhs]
	for i := lhs; i < rhs; i++ {
		if arr[i] < x {
			arr[i], arr[lhs] = arr[lhs], arr[i]
			lhs++
		}
	}
	arr[lhs], arr[rhs] = arr[rhs], arr[lhs]
	return lhs
}
