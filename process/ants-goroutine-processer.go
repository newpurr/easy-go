package process

import (
	"github.com/newpurr/easy-go/application"
	"github.com/panjf2000/ants/v2"
	"log"
)

type AntsGoroutineProcess struct {
	poolSize         int
	maxBlockingTasks int
}

func NewAntsGoroutineProcess(poolSize int, maxBlockingTasks int) *AntsGoroutineProcess {
	return &AntsGoroutineProcess{poolSize: poolSize, maxBlockingTasks: maxBlockingTasks}
}

func (s AntsGoroutineProcess) MustStart() {
	p, err := ants.NewPool(s.poolSize, ants.WithMaxBlockingTasks(s.maxBlockingTasks))
	if err != nil {
		panic(err)
	}
	application.AntsGoroutinePool = p
}

func (s AntsGoroutineProcess) Stop() {
	pool := application.AntsGoroutinePool
	if !pool.IsClosed() {
		pool.Release()
		log.Println("close the ants goroutine process")
	}
}
