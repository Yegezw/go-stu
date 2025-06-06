package main

import (
	"fmt"
)

// node 节点
type node[E comparable] struct {
	e    E
	next *node[E]
}

func newNode[E comparable](e E, next *node[E]) *node[E] {
	return &node[E]{e: e, next: next}
}

// LinkedList 递归单链表
type LinkedList[E comparable] struct {
	head *node[E]
	size int
}

// NewLinkedList 创建空链表
func NewLinkedList[E comparable]() *LinkedList[E] {
	return &LinkedList[E]{
		head: nil,
		size: 0,
	}
}

// GetSize 返回元素数量
func (list *LinkedList[E]) GetSize() int {
	return list.size
}

// IsEmpty 判断链表是否为空
func (list *LinkedList[E]) IsEmpty() bool {
	return list.size == 0
}

// Add 向指定位置插入元素
func (list *LinkedList[E]) Add(index int, e E) {
	if index < 0 || index > list.size {
		panic("index out of range, need 0 <= index <= size")
	}

	list.head = list.add(list.head, index, e)
	list.size++
}

// 以 node 为头结点的链表, 给它的 index 位置添加元素 e, 并返回新链表的头结点
func (list *LinkedList[E]) add(node *node[E], index int, e E) *node[E] {
	if index == 0 {
		return newNode(e, node)
	}

	node.next = list.add(node.next, index-1, e)
	return node
}

// AddFirst 向链表头部插入元素
func (list *LinkedList[E]) AddFirst(e E) {
	list.Add(0, e)
}

// AddLast 向链表尾部插入元素
func (list *LinkedList[E]) AddLast(e E) {
	list.Add(list.size, e)
}

// Remove 删除指定位置的元素
func (list *LinkedList[E]) Remove(index int) E {
	if index < 0 || index >= list.size {
		panic("index out of range, need 0 <= index < size")
	}

	head, e := list.remove(list.head, index)
	list.head = head
	list.size--
	return e
}

// 以 node 为头结点的链表, 删除它 index 位置的元素, 并返回新链表的头结点和被删除的元素 e
func (list *LinkedList[E]) remove(node *node[E], index int) (head *node[E], e E) {
	if index == 0 {
		return node.next, node.e
	}

	head, e = list.remove(node.next, index-1)
	node.next = head

	return node, e
}

// RemoveFirst 删除链表头部元素
func (list *LinkedList[E]) RemoveFirst() E {
	return list.Remove(0)
}

// RemoveLast 删除链表尾部元素
func (list *LinkedList[E]) RemoveLast() E {
	return list.Remove(list.size - 1)
}

// RemoveElement 删除链表中所有值为 e 的元素
func (list *LinkedList[E]) RemoveElement(e E) {
	list.head = list.removeElement(list.head, e)
}

// 以 node 为头结点的链表, 删除它中 "所有的" 元素 e, 并返回新链表的头结点
func (list *LinkedList[E]) removeElement(node *node[E], e E) (head *node[E]) {
	if node == nil {
		return nil
	}

	node.next = list.removeElement(node.next, e) // "所有的" 体现在这里

	if node.e == e {
		list.size--
		return node.next
	}
	return node
}

// Set 修改链表指定位置的元素
func (list *LinkedList[E]) Set(index int, e E) {

}

// 以 node 为头结点的链表, 把它 index 位置的元素修改为 e
func (list *LinkedList[E]) set(node *node[E], index int, e E) {
	if index == 0 {
		node.e = e
		return
	}
	list.set(node.next, index-1, e)
}

// Get 获取链表指定位置的元素
func (list *LinkedList[E]) Get(index int) E {
	if index < 0 || index >= list.size {
		panic("index out of range, need 0 <= index < size")
	}
	return list.get(list.head, index)
}

// 以 node 为头结点的链表, 获得它 index 位置的元素
func (list *LinkedList[E]) get(node *node[E], index int) E {
	if index == 0 {
		return node.e
	}
	return list.get(node.next, index-1)
}

// GetFirst 获取链表头部元素
func (list *LinkedList[E]) GetFirst() E {
	return list.Get(0)
}

// GetLast 获取链表尾部元素
func (list *LinkedList[E]) GetLast() E {
	return list.Get(list.size - 1)
}

// Contains 判断链表是否包含元素 e
func (list *LinkedList[E]) Contains(e E) bool {
	return list.contains(list.head, e)
}

// 以 node 为头结点的链表, 查询它中否有元素 e
func (list *LinkedList[E]) contains(node *node[E], e E) bool {
	if node == nil {
		return false
	}
	if node.e == e {
		return true
	}
	return list.contains(node.next, e)
}

// String 返回链表的字符串表示
func (list *LinkedList[E]) String() string {
	var sb string
	for node := list.head; node != nil; node = node.next {
		sb += fmt.Sprintf("%v->", node.e)
	}
	sb += "nil"
	return sb
}
