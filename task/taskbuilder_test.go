package task_test

import (
	"github.com/arielsrv/golang-toolkit/task"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestBuilder_ForkJoin(t *testing.T) {
	var future1, future2 *task.Task[int]

	tb := &task.Builder{}

	start := time.Now()
	tb.ForkJoin(func(c *task.Awaitable) {
		future1 = task.Await[int](c, func() (int, error) {
			time.Sleep(time.Millisecond * 1000)
			return 2, nil
		})
		future2 = task.Await[int](c, func() (int, error) {
			time.Sleep(time.Millisecond * 1000)
			return 3, nil
		})
	})

	assert.NotNil(t, future1.Result)
	assert.NoError(t, future1.Err)
	assert.Equal(t, 2, future1.Result)

	assert.NotNil(t, future2.Result)
	assert.NoError(t, future2.Err)
	assert.Equal(t, 3, future2.Result)

	assert.Equal(t, 5, future1.Result+future2.Result)

	end := time.Since(start)
	log.Println(end)

	assert.Greater(t, time.Millisecond*(1000*1.01), end)
}
