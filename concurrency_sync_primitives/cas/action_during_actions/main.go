package main

import "sync/atomic"

type Data struct {
	count atomic.Int32
}

func (d *Data) Process() {
	d.count.Add(1)
	if d.count.CompareAndSwap(100, 0) {
		// do something...
	}
}
