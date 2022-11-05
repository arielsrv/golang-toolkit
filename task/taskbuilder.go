package task

import (
	"container/list"
	"github.com/alitto/pond"
)

type Task[T any] struct {
	Result T
	Err    error
}

type Awaitable struct {
	list        list.List
	pool        *pond.WorkerPool
	taskBuilder *Builder
}

func Await[T any](c *Awaitable, f func() (T, error)) *Task[T] {
	fr := new(Task[T])

	future := func() {
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
	MaxWorkers  int
	MaxCapacity int
}

func (tb *Builder) ForkJoin(f func(*Awaitable)) {
	c := new(Awaitable)
	c.taskBuilder = tb

	f(c)

	c.pool = pond.New(tb.MaxWorkers, tb.MaxCapacity)

	for e := c.list.Front(); e != nil; e = e.Next() {
		action := e.Value.(func())
		c.pool.Submit(func() {
			action()
		})
	}

	c.pool.StopAndWait()
}
