package Week_01

import (
	"fmt"
	"sync"

	"github.com/dingkegithub/algorithm019/utils"
)

var (
	ErrNoSpace = fmt.Errorf("no space to store element")
	ErrBlank   = fmt.Errorf("blank container")
)

type Dqueue struct {
	mutex    sync.RWMutex
	capacity int
	size     int
	compare  utils.Compare
	data     []interface{}
}

func NewDqueue(opts ...utils.ParamOptionFunc) *Dqueue {
	dq := &Dqueue{
		size:     0,
		capacity: 10,
		compare:  utils.CompareNum,
	}

	dq.Apply(opts...)
	dq.data = make([]interface{}, dq.capacity)

	return dq
}

func WithDqueueCapacity(c int) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		dq := target.(*Dqueue)
		dq.capacity = c
	}
}

func WithDqueueCompare(c utils.Compare) utils.ParamOptionFunc {
	return func(target utils.ParamOption) {
		dq := target.(*Dqueue)
		dq.compare = c
	}
}

func (dq *Dqueue) Apply(fs ...utils.ParamOptionFunc) {
	for _, f := range fs {
		f(dq)
	}
}

func (dq *Dqueue) addIdx(idx int, e interface{}) {
	for i := dq.size; i > idx; i-- {
		dq.data[i] = dq.data[i-1]
	}
	dq.data[idx] = e
	dq.size += 1
}

func (dq *Dqueue) delIdx(idx int) interface{} {
	d := dq.data[idx]

	for i := idx; i < dq.size-1; i++ {
		dq.data[i] = dq.data[i+1]
	}
	dq.size -= 1

	return d
}

// Inserts the specified element at the front of this deque if it is possible to do so immediately
// without violating capacity restrictions, return an ErrNoSpace if no space is currently available.
func (dq *Dqueue) AddFirst(e interface{}) error {
	if dq.IsFull() {
		return ErrNoSpace
	}

	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.addIdx(0, e)
	return nil
}

// Inserts the specified element at the end of this deque if it is possible to do so immediately without
// violating capacity restrictions, return an ErrNoSpace if no space is currently available.
func (dq *Dqueue) AddLast(e interface{}) error {
	if dq.IsFull() {
		return ErrNoSpace
	}

	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.addIdx(dq.size, e)
	return nil
}

// Retrieves and removes the first element of this deque.
func (dq *Dqueue) RemoveFirst() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, ErrBlank
	}

	dq.mutex.Lock()
	defer dq.mutex.Unlock()
	return dq.delIdx(0), nil
}

// Retrieves and removes the last element of this deque.
func (dq *Dqueue) RemoveLast() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, ErrBlank
	}

	dq.mutex.Lock()
	defer dq.mutex.Unlock()
	return dq.delIdx(dq.size - 1), nil
}

// Retrieves and removes the first element of this deque, or returns nil if this deque is empty
func (dq *Dqueue) PollFirst() interface{} {
	if dq.IsEmpty() {
		return nil
	}

	dq.mutex.Lock()
	defer dq.mutex.Unlock()
	return dq.delIdx(0)
}

// Retrieves and removes the last element of this deque, or returns nil if this deque is empty.
func (dq *Dqueue) PollLast() interface{} {
	if dq.IsEmpty() {
		return nil
	}

	dq.mutex.Lock()
	defer dq.mutex.Unlock()
	return dq.delIdx(dq.size - 1)
}

// Retrieves, but does not remove, the first element of this deque.
func (dq *Dqueue) GetFirst() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, ErrBlank
	}

	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.data[0], nil
}

// Retrieves, but does not remove, the last element of this deque.
func (dq *Dqueue) GetLast() (interface{}, error) {
	if dq.IsEmpty() {
		return nil, ErrBlank
	}

	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.data[dq.size-1], nil
}

// Retrieves, but does not remove, the first element of this deque, or returns nil if this deque is empty
func (dq *Dqueue) PeekFirst() interface{} {
	if dq.IsEmpty() {
		return nil
	}

	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.data[0]
}

// Retrieves, but does not remove, the last element of this deque, or returns nil if this deque is empty.
func (dq *Dqueue) PeekLast() interface{} {
	if dq.IsEmpty() {
		return nil
	}

	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.data[dq.size-1]
}

func (dq *Dqueue) Size() int {
	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.size
}

func (dq *Dqueue) Cap() int {
	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.capacity
}

func (dq *Dqueue) IsFull() bool {
	dq.mutex.RLock()
	defer dq.mutex.RUnlock()

	return dq.size >= dq.capacity
}

func (dq *Dqueue) IsEmpty() bool {
	dq.mutex.RLock()
	defer dq.mutex.RUnlock()
	return dq.size == 0
}
