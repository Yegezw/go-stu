package main

func swap[E any](arr []E, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func eqInt(a, b int) bool {
	return a == b
}

func eqString(a, b string) bool {
	return a == b
}

func lt(a, b int) bool {
	return a < b
}

func gt(a, b int) bool {
	return a > b
}

func le(a, b int) bool {
	return a <= b
}

func ge(a, b int) bool {
	return a >= b
}

// [0 ... max] 循环 + 1
func addOne(cur int, max int) int {
	if cur == max {
		return 0
	}
	return cur + 1
}

// [0 ... max] 循环 - 1
func subOne(cur int, max int) int {
	if cur == 0 {
		return max
	}
	return cur - 1
}
