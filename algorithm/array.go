package main

import (
	"fmt"
)

// Array 动态数组
type Array[E comparable] struct {
	data     []E
	size     int
	modCount int
}

// NewArray 创建指定容量的数组
func NewArray[E comparable](capacity int) *Array[E] {
	return &Array[E]{
		data:     make([]E, capacity),
		size:     0,
		modCount: 0,
	}
}

// NewEmptyArray 创建默认容量的空数组
func NewEmptyArray[E comparable]() *Array[E] {
	return NewArray[E](10)
}

// NewArrayFromSlice 从切片创建数组
func NewArrayFromSlice[E comparable](slice []E) *Array[E] {
	length := len(slice)
	data := make([]E, length)
	copy(data, slice)
	return &Array[E]{
		data:     data,
		size:     length,
		modCount: 0,
	}
}

// GetSize 返回元素数量
func (array *Array[E]) GetSize() int {
	return array.size
}

// GetCapacity 返回当前容量
func (array *Array[E]) GetCapacity() int {
	return cap(array.data)
}

// IsEmpty 判断数组是否为空
func (array *Array[E]) IsEmpty() bool {
	return array.size == 0
}

// Add 在指定位置插入元素
func (array *Array[E]) Add(index int, e E) {
	if index < 0 || index > array.size {
		panic("index out of range, need 0 <= index <= size")
	}
	if array.size == array.GetCapacity() {
		array.resize(array.GetCapacity() * 2)
	}

	array.modCount++
	copy(array.data[index+1:], array.data[index:array.size])
	array.data[index] = e
	array.size++
}

// AddFirst 在头部插入元素
func (array *Array[E]) AddFirst(e E) {
	array.Add(0, e)
}

// AddLast 在尾部插入元素
func (array *Array[E]) AddLast(e E) {
	array.Add(array.size, e)
}

// Remove 删除指定位置元素
func (array *Array[E]) Remove(index int) E {
	var empty E
	if index < 0 || index >= array.size {
		panic("index out of range, need 0 <= index < size")
	}

	array.modCount++
	ret := array.data[index]
	copy(array.data[index:], array.data[index+1:array.size])
	array.size--
	array.data[array.size] = empty // help gc

	if array.size/2 != 0 && array.size == array.GetCapacity()/4 {
		array.resize(array.GetCapacity() / 2)
	}
	return ret
}

// RemoveFirst 删除头部元素
func (array *Array[E]) RemoveFirst() E {
	return array.Remove(0)
}

// RemoveLast 删除尾部元素
func (array *Array[E]) RemoveLast() E {
	return array.Remove(array.size - 1)
}

// RemoveElement 删除指定元素
func (array *Array[E]) RemoveElement(e E) {
	index := array.Find(e)
	if index != -1 {
		array.Remove(index)
	}
}

// Set 修改指定位置元素
func (array *Array[E]) Set(index int, e E) {
	if index < 0 || index >= array.size {
		panic("index out of range, need 0 <= index < size")
	}
	array.modCount++
	array.data[index] = e
}

// Get 获取指定位置元素
func (array *Array[E]) Get(index int) E {
	if index < 0 || index >= array.size {
		panic("index out of range, need 0 <= index < size")
	}
	return array.data[index]
}

// GetFirst 获取第一个元素
func (array *Array[E]) GetFirst() E {
	return array.Get(0)
}

// GetLast 获取最后一个元素
func (array *Array[E]) GetLast() E {
	return array.Get(array.size - 1)
}

// Contains 判断元素是否存在
func (array *Array[E]) Contains(e E) bool {
	for i := 0; i < array.size; i++ {
		if array.data[i] == e {
			return true
		}
	}
	return false
}

// Find 查找元素位置
func (array *Array[E]) Find(e E) int {
	for i := 0; i < array.size; i++ {
		if array.data[i] == e {
			return i
		}
	}
	return -1
}

// resize 调整数组容量
func (array *Array[E]) resize(newCapacity int) {
	newData := make([]E, newCapacity)
	copy(newData, array.data[:array.size])
	array.data = newData
}

// ToSlice 将动态数组转换为切片
func (array *Array[E]) ToSlice() []E {
	return append([]E{}, array.data[:array.size]...)
}

// Swap 交换两个位置的元素
func (array *Array[E]) Swap(i, j int) {
	if i < 0 || i >= array.size || j < 0 || j >= array.size {
		panic("swap failed, indices out of range")
	}
	array.modCount++
	array.data[i], array.data[j] = array.data[j], array.data[i]
}

// String 返回数组的字符串表示
func (array *Array[E]) String() string {
	sb := fmt.Sprintf("Array: size = %d, capacity = %d\n", array.size, array.GetCapacity())
	sb += "["
	for i := 0; i < array.size; i++ {
		sb += fmt.Sprintf("%v", array.data[i])
		if i != array.size-1 {
			sb += ", "
		}
	}
	sb += "]"
	return sb
}

// Iterator 返回迭代器
func (array *Array[E]) Iterator() *ArrayIterator[E] {
	return &ArrayIterator[E]{
		array:            array,
		cursor:           0,
		lastRet:          -1,
		expectedModCount: array.modCount,
	}
}

// ArrayIterator 动态数组迭代器
type ArrayIterator[E comparable] struct {
	array            *Array[E]
	cursor           int // 待遍历
	lastRet          int // 已遍历
	expectedModCount int
}

// HasNext 判断是否有下一个元素
func (it *ArrayIterator[E]) HasNext() bool {
	return it.cursor != it.array.size
}

// Next 返回下一个元素
func (it *ArrayIterator[E]) Next() E {
	if it.expectedModCount != it.array.modCount {
		panic("concurrent modification detected")
	}
	if it.cursor >= it.array.size {
		panic("no more elements")
	}
	it.lastRet = it.cursor
	it.cursor++
	return it.array.data[it.lastRet]
}

// Remove 删除当前元素
func (it *ArrayIterator[E]) Remove() {
	if it.lastRet < 0 {
		panic("illegal state, call Next() first")
	}
	if it.expectedModCount != it.array.modCount {
		panic("concurrent modification detected")
	}

	it.array.Remove(it.lastRet)
	it.cursor = it.lastRet
	it.lastRet = -1
	it.expectedModCount = it.array.modCount
}
