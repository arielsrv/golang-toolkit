package task_test

import (
	"github.com/arielsrv/golang-toolkit/task"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// https://go.dev/play/p/yViud-GNlh2
func TestBuilder_ForkJoin(t *testing.T) {
	var future1, future2 *task.Task

	tb := &task.Builder{}

	start := time.Now()

	tb.ForkJoin(func(c *task.Awaitable) {
		future1 = task.Await[int](c, func() (int, error) {
			time.Sleep(time.Millisecond * 1000)
			return 1, nil
		})
		future2 = task.Await[int](c, func() (int, error) {
			time.Sleep(time.Millisecond * 1000)
			return 1, nil
		})
	})

	assert.NoError(t, future1.Err)
	assert.NoError(t, future2.Err)

	actual1 := task.Result[int](future1)
	actual2 := task.Result[int](future2)

	assert.Equal(t, 2, *actual1+*actual2)
	assert.Greater(t, time.Millisecond*(1000*1.01), time.Since(start))
}
