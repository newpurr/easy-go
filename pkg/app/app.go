package app

import (
	"context"
	"github.com/newpurr/easy-go/pkg/signalc"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
	"syscall"
)

type App struct {
	eg         errgroup.Group
	l          sync.Mutex
	p          []Processer
	stopSignal chan struct{}
}

func NewApp() *App {
	return &App{
		eg:         errgroup.Group{},
		l:          sync.Mutex{},
		p:          []Processer{},
		stopSignal: make(chan struct{}),
	}
}

func (a *App) WithProcesser(p ...Processer) *App {
	a.l.Lock()
	defer a.l.Unlock()
	for _, processer := range p {
		a.p = append(a.p, processer)
	}

	return a
}

func (a *App) Start(ctx context.Context) {
	a.l.Lock()
	defer a.l.Unlock()

	go func() {
		select {
		case <-signalc.Wait(syscall.SIGINT, syscall.SIGTERM):
			close(a.stopSignal)
		case <-ctx.Done():
			close(a.stopSignal)
		}
	}()
	for _, processer := range a.p {
		p := processer
		a.eg.Go(func() error {
			go p.MustStart()
			<-a.stopSignal
			p.Stop()
			return nil
		})
	}

	_ = a.eg.Wait()

	log.Println("server has exited.")
}

func FastStart(p ...Processer) {
	starter := NewApp()
	starter.WithProcesser(p...)
	starter.Start(context.Background())
}
