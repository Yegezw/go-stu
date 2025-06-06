package main

import "fmt"

// Dequeue 双端队列
type Dequeue[E any] struct {
	// data[front ... tail)
	data  []E
	front int
	tail  int
	size  int
}

// NewDequeue 创建指定容量的双端队列
func NewDequeue[E any](capacity int) *Dequeue[E] {
	return &Dequeue[E]{
		data:  make([]E, capacity),
		front: 0,
		tail:  0,
		size:  0,
	}
}

// NewDefaultDequeue 创建默认容量的双端队列
func NewDefaultDequeue[E any]() *Dequeue[E] {
	return NewDequeue[E](10)
}

// AddFront 向队头添加元素
func (dequeue *Dequeue[E]) AddFront(e E) {
	if dequeue.size == dequeue.GetCapacity() {
		dequeue.resize(dequeue.GetCapacity() * 2)
	}

	dequeue.front = subOne(dequeue.front, dequeue.GetCapacity()-1)
	dequeue.data[dequeue.front] = e
	dequeue.size++
}

// AddLast 向队尾添加元素
func (dequeue *Dequeue[E]) AddLast(e E) {
	if dequeue.size == dequeue.GetCapacity() {
		dequeue.resize(dequeue.GetCapacity() * 2)
	}

	dequeue.data[dequeue.tail] = e
	dequeue.tail = addOne(dequeue.tail, dequeue.GetCapacity()-1)
	dequeue.size++
}

// RemoveFront 从队头移除元素
func (dequeue *Dequeue[E]) RemoveFront() E {
	var empty E
	if dequeue.IsEmpty() {
		panic("dequeue is empty")
	}

	ret := dequeue.data[dequeue.front]
	dequeue.data[dequeue.front] = empty // help gc
	dequeue.front = addOne(dequeue.front, dequeue.GetCapacity()-1)
	dequeue.size--

	if dequeue.size/2 != 0 && dequeue.size == dequeue.GetCapacity()/4 {
		dequeue.resize(dequeue.GetCapacity() / 2)
	}
	return ret
}

// RemoveLast 从队尾移除元素
func (dequeue *Dequeue[E]) RemoveLast() E {
	var empty E
	if dequeue.IsEmpty() {
		panic("dequeue is empty")
	}

	dequeue.tail = subOne(dequeue.tail, dequeue.GetCapacity()-1)
	ret := dequeue.data[dequeue.tail]
	dequeue.data[dequeue.tail] = empty // help gc
	dequeue.size--

	if dequeue.size/2 != 0 && dequeue.size == dequeue.GetCapacity()/4 {
		dequeue.resize(dequeue.GetCapacity() / 2)
	}
	return ret
}

// GetFront 获取队头元素
func (dequeue *Dequeue[E]) GetFront() E {
	if dequeue.IsEmpty() {
		panic("dequeue is empty")
	}
	return dequeue.data[dequeue.front]
}

// GetLast 获取队尾元素
func (dequeue *Dequeue[E]) GetLast() E {
	if dequeue.IsEmpty() {
		panic("dequeue is empty")
	}
	index := subOne(dequeue.tail, dequeue.GetCapacity()-1)
	return dequeue.data[index]
}

// IsEmpty 判断队列是否为空
func (dequeue *Dequeue[E]) IsEmpty() bool {
	return dequeue.size == 0
}

// GetSize 返回元素数量
func (dequeue *Dequeue[E]) GetSize() int {
	return dequeue.size
}

// GetCapacity 返回当前容量
func (dequeue *Dequeue[E]) GetCapacity() int {
	return cap(dequeue.data)
}

func (dequeue *Dequeue[E]) resize(newCapacity int) {
	newData := make([]E, newCapacity)
	for i := 0; i < dequeue.size; i++ {
		newData[i] = dequeue.data[(dequeue.front+i)%dequeue.GetCapacity()]
	}
	dequeue.data = newData
	dequeue.front = 0
	dequeue.tail = dequeue.size
}

// String 返回队列的字符串表示
func (dequeue *Dequeue[E]) String() string {
	sb := fmt.Sprintf("Deque: size = %d, capacity = %d\n", dequeue.size, dequeue.GetCapacity())
	sb += "["
	for i := 0; i < dequeue.size; i++ {
		idx := (dequeue.front + i) % dequeue.GetCapacity()
		sb += fmt.Sprintf("%v", dequeue.data[idx])
		if i != dequeue.size-1 {
			sb += ", "
		}
	}
	sb += "]"
	return sb
}
