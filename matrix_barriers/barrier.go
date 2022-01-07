package main

import "sync"

type Barrier struct {
	total int
	count int
	mx    *sync.Mutex
	cv    *sync.Cond
}

func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{size, size, lockToUse, condToUse}
}

func (b *Barrier) Wait() {
	b.mx.Lock()
	b.count -= 1
	if b.count == 0 {
		b.count = b.total
		b.cv.Broadcast()
	} else {
		b.cv.Wait()
	}
	b.mx.Unlock()
}
