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
