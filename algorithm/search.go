package main

// LinearSearch 线性查找
func LinearSearch[E any](data []E, target E, equals func(a, b E) bool) (index int) {
	for i := range data {
		if equals(data[i], target) {
			return i
		}
	}
	return -1
}

// BinarySearch 二分查找
func BinarySearch[E any](data []E, target E, le func(a, b E) bool, gt func(a, b E) bool) (index int) {
	l := 0
	r := len(data) - 1
	var mid int

	for l <= r {
		mid = l + (r-l)/2
		if le(target, data[mid]) {
			r = mid - 1
		} else if gt(target, data[mid]) {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
