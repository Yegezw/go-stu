package main

// List 表示一个 "可以保存任何类型值" 的单链表
type List[T any] struct {
	val  T
	next *List[T]
}
