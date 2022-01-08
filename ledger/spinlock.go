package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type Spinlock int32

func (s *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *Spinlock) Unlock() {
	atomic.StoreInt32((*int32)(s), 0)
}

func NewSpinlock() sync.Locker {
	var lock Spinlock
	return &lock
}
