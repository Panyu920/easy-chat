package lock

import (
	"runtime"
	"sync/atomic"
)

type SpinLockType int32

type SpinLock struct {
	value SpinLockType
}

func NewSpinLock() *SpinLock {
	return &SpinLock{}
}

func (s *SpinLock) Lock() {
	tryCount := 0
	for !atomic.CompareAndSwapInt32((*int32)(&s.value), 0, 1) {
		if tryCount < 4 {
			tryCount++
		} else {
			runtime.Gosched()
		}
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(&s.value), 0)
}
