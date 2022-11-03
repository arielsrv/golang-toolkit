package task_test

import (
	"github.com/arielsrv/golang-toolkit/task"
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://go.dev/play/p/yViud-GNlh2
func TestBuilder_ForkJoin(t *testing.T) {
	var future1, future2 *task.Task

	tb := &task.Builder{}

	tb.ForkJoin(func(c *task.Awaitable) {
		future1 = task.Await[int](c, func() (int, error) { return 1, nil })
		future2 = task.Await[int](c, func() (int, error) { return 1, nil })
	})

	assert.NoError(t, future1.Err)
	assert.NoError(t, future2.Err)

	actual1 := (*int)(future1.Ptr)
	actual2 := (*int)(future1.Ptr)

	assert.Equal(t, 2, *actual1+*actual2)
}
