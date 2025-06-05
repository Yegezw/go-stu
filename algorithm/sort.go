package main

import (
	"math/rand"
)

// SelectionSort 选择排序
func SelectionSort[E any](arr []E, lt func(a, b E) bool) {
	length := len(arr)

	for i := 0; i < length; i++ {
		minIndex := i

		for j := i + 1; j < length; j++ {
			if lt(arr[j], arr[minIndex]) {
				minIndex = j
			}
		}

		swap(arr, i, minIndex)
	}
}

// InsertionSort 插入排序
func InsertionSort[E any](arr []E, gt func(a, b E) bool) {
	length := len(arr)

	for i := 1; i < length; i++ {
		k := arr[i]

		var j int
		for j = i; j-1 >= 0 && gt(arr[j-1], k); j-- {
			arr[j] = arr[j-1]
		}

		arr[j] = k
	}
}

// MergerSort 归并排序
func MergerSort[E any](arr []E, le func(a, b E) bool) {
	length := len(arr)
	temp := make([]E, length)
	mergerSort(arr, 0, length-1, temp, le)
}

func mergerSort[E any](arr []E, l int, r int, temp []E, le func(a, b E) bool) {
	if l >= r {
		return
	}

	mid := l + (r-l)/2
	mergerSort[E](arr, l, mid, temp, le)   // arr[l, mid]
	mergerSort[E](arr, mid+1, r, temp, le) // arr[mid + 1, r]

	if !le(arr[mid], arr[mid+1]) {
		merge(arr, l, mid, r, temp, le)
	}
}

// merge
// 合并两个有序数组 arr[l, mid] 和 arr[mid + 1, r], 使得 arr[l, r] 整体有序
func merge[E any](arr []E, l int, mid int, r int, temp []E, le func(a, b E) bool) {
	copy(temp[l:r+1], arr[l:r+1])

	p1 := l       // temp[l, mid]
	p2 := mid + 1 // temp[mid + 1, r]
	i := l        // arr[l, r]

	for p1 <= mid && p2 <= r {
		if le(temp[p1], temp[p2]) {
			arr[i] = temp[p1]
			i++
			p1++
		} else {
			arr[i] = temp[p2]
			i++
			p2++
		}
	}

	for p1 <= mid {
		arr[i] = temp[p1]
		i++
		p1++
	}

	for p2 <= r {
		arr[i] = temp[p2]
		i++
		p2++
	}
}

// QuickSort 快速排序
func QuickSort[E any](arr []E, lt func(a, b E) bool, gt func(a, b E) bool) {
	quickSort(arr, 0, len(arr)-1, lt, gt)
}

func quickSort[E any](arr []E, l int, r int, lt func(a, b E) bool, gt func(a, b E) bool) {
	if l >= r {
		return
	}

	p := partition(arr, l, r, lt, gt)

	quickSort[E](arr, l, p-1, lt, gt)
	quickSort[E](arr, p+1, r, lt, gt)
}

func partition[E any](arr []E, l int, r int, lt func(a, b E) bool, gt func(a, b E) bool) int {
	p := rand.Intn(r-l) + l
	swap(arr, l, p)

	v := arr[l]
	p1 := l + 1
	p2 := r

	for {
		for p1 <= p2 && lt(arr[p1], v) {
			p1++
		}
		for p1 <= p2 && gt(arr[p2], v) {
			p2--
		}

		if p1 >= p2 {
			break
		}

		swap(arr, p1, p2)
		p1++
		p2--
	}

	swap(arr, l, p2)
	return p2
}
