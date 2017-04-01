package atomicCounter

import "sync/atomic"

type Counter struct {
	count uint64
}

func (counter *Counter) Incr() {
	atomic.AddUint64(&counter.count, 1)
}

func (counter *Counter) Decr() {
	atomic.AddUint64(&counter.count, ^uint64(0))
}

func (counter *Counter) Get() uint64 {
	return atomic.LoadUint64(&counter.count)
}
