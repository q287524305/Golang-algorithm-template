package main

import "fmt"

func mergeSort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)

	var k, i, j = 0, l, mid+1
	tmp := make([]int, r-l+1)
	for ; i <= mid && j <= r; k++ {
		if q[i] < q[j] {
			tmp[k] = q[i]
			i++
		} else {
			tmp[k] = q[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp[k] = q[i]
		k++
	}
	for ; j <= r; j++ {
		tmp[k] = q[j]
		k++
	}
	copy(q[l:r+1], tmp)
}

func main() {
	a := []int{5, 2, 3, 7, 6, 9, 0, 8, 4, 1}
	mergeSort(a, 0, len(a)-1)
	fmt.Print(a)
}
