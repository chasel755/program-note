package data-struct

/**
 * 循环队列
 */
type MyCircularQueue struct {
	size int
	head int
	tail int
	data []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue {
		size: k,
		head: -1,
		tail: -1,
		data: make([]int, k),
	}
}

func (this MyCircularQueue) IsEmpty() bool {
	if this.head == -1 {
		return true
	}
	return false
}

func (this MyCircularQueue) IsFull() bool {
	if (this.tail + 1) % this.size == this.head {
		return true
	}
	return false
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % this.size
	this.data[this.tail] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	this.head = (this.head + 1) % this.size
	return true
}

func (this MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]	
}

func (this MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]	
}







