package main

import "fmt"

func main() {
	testSearch()
	testSort()
}

func testSearch() {
	arr1 := []int{1, 2, 3}
	fmt.Println(
		LinearSearch(arr1, 10, eqInt),
	)

	arr2 := []string{"a", "b", "c", "d"}
	fmt.Println(
		LinearSearch(arr2, "c", eqString),
	)

	arr3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(
		BinarySearch(arr3, 4, lt, gt),
	)
}

func testSort() {
	arr := []int{1, 4, 2, 5, 9, 6, 8, 3, 7}

	SelectionSort(arr, lt) // 选择排序
	InsertionSort(arr, gt) // 插入排序
	MergerSort(arr, le)    // 归并排序
	QuickSort(arr, lt, gt) // 快速排序

	fmt.Println(arr)
}
