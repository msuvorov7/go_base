package sorting

func merge(arr []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	lhs := make([]int, n1)
	rhs := make([]int, n2)
	for i := 0; i < n1; i++ {
		lhs[i] = arr[p+i]
	}
	for i := 0; i < n2; i++ {
		rhs[i] = arr[q+i+1]
	}
	i := 0
	j := 0
	k := p
	for i < n1 && j < n2 {
		if lhs[i] <= rhs[j] {
			arr[k] = lhs[i]
			i++
		} else {
			arr[k] = rhs[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = lhs[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = rhs[j]
		j++
		k++
	}
}

func mergeSort(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(arr, p, q)
		mergeSort(arr, q+1, r)
		merge(arr, p, q, r)
	}
}

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}
