package main

import "fmt"

func main() {
	testSearch()
	fmt.Println()

	testSort()
	fmt.Println()

	testArray()
	fmt.Println()

	testDequeue()
	fmt.Println()
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

func testArray() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	array1 := NewArrayFromSlice([]int{1, 2, 3})
	array2 := NewArrayFromSlice([]int{4, 5, 6})
	slice1 := array1.ToSlice()
	slice2 := array2.ToSlice()

	array := NewEmptyArray[int]()
	for _, v := range slice1 {
		array.AddLast(v)
	}
	for _, v := range slice2 {
		array.AddFirst(v)
	}
	fmt.Println(array)

	it := array.Iterator()
	for it.HasNext() {
		e := it.Next()
		// 删除偶数
		if (e & 1) == 0 {
			it.Remove()
			continue
		}
	}
	fmt.Println(array)
}

func testDequeue() {
	dequeue := NewDefaultDequeue[int]()
	for i := 0; i < 16; i++ {
		if (i & 1) == 0 {
			dequeue.AddLast(i)
		} else {
			dequeue.AddFront(i)
		}
		fmt.Println(dequeue)
	}

	for i := 0; !dequeue.IsEmpty(); i++ {
		if (i & 1) == 0 {
			dequeue.RemoveFront()
		} else {
			dequeue.RemoveLast()
		}
		fmt.Println(dequeue)
	}
}
