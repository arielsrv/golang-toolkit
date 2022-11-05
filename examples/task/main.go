package main

import (
	"github.com/arielsrv/golang-toolkit/task"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var future1, future2, future3 *task.Task[int]

	tb := &task.Builder{
		MaxWorkers: runtime.NumCPU() - 1,
	}

	start := time.Now()
	tb.ForkJoin(func(c *task.Awaitable) {
		future1 = task.Await[int](c, GetNumber)
		future2 = task.Await[int](c, GetNumber)
		future3 = task.Await[int](c, GetNumber)
	})

	log.Println(future1.Result)
	log.Println(future2.Result)
	log.Println(future3.Result)

	end := time.Since(start)
	log.Println(end)
}

func GetNumber() (int, error) {
	value := rand.Int()
	time.Sleep(time.Millisecond * 1000)
	log.Println("done ...")
	return value, nil
}
