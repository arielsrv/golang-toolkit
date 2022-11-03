package task

import (
	"container/list"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Task struct {
	Ptr unsafe.Pointer
	Err error
}

type Concurrent struct {
	list        list.List
	wg          sync.WaitGroup
	taskBuilder *Builder
}

func Await[T any](c *Concurrent, f func() (T, error)) *Task {
	fr := new(Task)

	future := func() {
		defer c.wg.Done()
		r, err := f()
		atomic.StorePointer(&fr.Ptr, unsafe.Pointer(&r))
		fr.Err = err
	}

	c.list.PushBack(future)

	return fr
}

type Builder struct {
}

func (tb *Builder) ForkJoin(f func(*Concurrent)) {
	c := new(Concurrent)
	c.taskBuilder = tb

	f(c)

	c.wg.Add(c.list.Len())

	for e := c.list.Front(); e != nil; e = e.Next() {
		go e.Value.(func())()
	}

	c.wg.Wait()
}
