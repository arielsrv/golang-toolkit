package task

import (
	"container/list"
	"sync"
)

type Task[T any] struct {
	Result T
	Err    error
}

type Awaitable struct {
	list        list.List
	wg          sync.WaitGroup
	taskBuilder *Builder
}

func Await[T any](c *Awaitable, f func() (T, error)) *Task[T] {
	fr := new(Task[T])

	future := func() {
		defer c.wg.Done()
		r, err := f()
		fr.Result = r
		fr.Err = err
	}

	c.list.PushBack(future)

	return fr
}

type Result[T any] struct {
	Result T
	Err    error
}

type Builder struct {
}

func (tb *Builder) ForkJoin(f func(*Awaitable)) {
	c := new(Awaitable)
	c.taskBuilder = tb

	f(c)

	c.wg.Add(c.list.Len())

	for e := c.list.Front(); e != nil; e = e.Next() {
		go e.Value.(func())()
	}

	c.wg.Wait()
}
