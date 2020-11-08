package Week_01

import "sync"

type MyCircularDeque struct {
	front    int
	rear     int
	capacity int
	data     []int
	mutex    sync.RWMutex
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		front:    0,
		rear:     0,
		capacity: k + 1,
		data:     make([]int, k+1, k+1),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.front = (this.front - 1 + this.capacity) % this.capacity
	this.data[this.front] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % this.capacity
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.front = (this.front + 1) % this.capacity
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.rear = (this.rear - 1 + this.capacity) % this.capacity
	return true

}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.data[this.front]

}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.data[(this.rear-1+this.capacity)%this.capacity]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.rear == this.front
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return (this.rear+1)%this.capacity == this.front
}
