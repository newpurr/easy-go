package app

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"time"
)

type AntsMnsGoroutineProcess struct {
	p         *ants.Pool
	closeChan chan struct{}
}

func NewAntsMnsGoroutineProcess() *AntsMnsGoroutineProcess {
	return &AntsMnsGoroutineProcess{
		closeChan: make(chan struct{}),
	}
}

func (s *AntsMnsGoroutineProcess) MustStart() {
	p, err := ants.NewPool(5, ants.WithMaxBlockingTasks(5))
	if err != nil {
		panic(err)
	}
	s.p = p
	for i := 0; i < 5; i++ {
		_ = p.Submit(func() {
			for {
				select {
				case <-s.closeChan:
					break
				default:
				}

				fmt.Println("hello world")
				time.Sleep(5 * time.Second)
			}
		})
	}
}

func (s *AntsMnsGoroutineProcess) Stop() {
	pool := s.p
	if !pool.IsClosed() {
		close(s.closeChan)
		pool.Release()
		log.Println("close the ants goroutine process")
	}
}
