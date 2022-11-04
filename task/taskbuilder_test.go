package task_test

import (
	"github.com/arielsrv/golang-toolkit/task"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// https://go.dev/play/p/yViud-GNlh2
func TestBuilder_ForkJoin(t *testing.T) {
	var future1 *task.Task[int]
	var future2 *task.Task[int]

	tb := &task.Builder{}

	start := time.Now()

	tb.ForkJoin(func(c *task.Awaitable) {
		future1 = task.Await[int](c, func() (int, error) {
			log.Println("future1")
			time.Sleep(time.Millisecond * 1000)
			return 1, nil
		})
		future2 = task.Await[int](c, func() (int, error) {
			log.Println("future2")
			time.Sleep(time.Millisecond * 1000)
			return 1, nil
		})
	})

	actual1 := future1.GetResult()
	assert.NotNil(t, actual1.Result)
	assert.NoError(t, actual1.Err)
	assert.Equal(t, 1, actual1.Result)

	actual2 := future2.GetResult()
	assert.NotNil(t, actual2.Result)
	assert.NoError(t, actual2.Err)
	assert.Equal(t, 1, actual2.Result)

	assert.Greater(t, time.Millisecond*(1000*1.01), time.Since(start))
}
