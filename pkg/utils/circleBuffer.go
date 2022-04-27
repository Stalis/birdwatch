package utils

import "sync"

type CircleBuffer struct {
	items     []interface{}
	capacity  int
	nextIndex int
	closed    bool
	mutex     sync.RWMutex
}

func NewCircleBuffer(capacity int) *CircleBuffer {
	return &CircleBuffer{
		items:     make([]interface{}, capacity),
		capacity:  capacity,
		nextIndex: 0,
	}
}

func (b *CircleBuffer) Add(val interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	index := b.nextIndex
	if b.nextIndex >= b.capacity {
		index -= b.capacity
	}

	b.items[index] = val

	b.nextIndex = index + 1
	if b.nextIndex == b.capacity {
		b.closed = true
	}
}

func (b *CircleBuffer) Get(index int) interface{} {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return b.items[index]
}

func (b *CircleBuffer) Closed() bool {
	return b.closed
}

func (b *CircleBuffer) Length() int {
	if b.closed {
		return b.capacity
	}

	return b.capacity - b.nextIndex
}

func (b *CircleBuffer) Items() []interface{} {
	return b.items
}
