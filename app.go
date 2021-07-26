package easygo

import (
	"context"
	"github.com/newpurr/easy-go/boot"
	"github.com/newpurr/easy-go/event"
	"github.com/newpurr/easy-go/process"
	"github.com/newpurr/easy-go/signalc"
	"golang.org/x/sync/errgroup"
	"log"
	"syscall"
)

var (
	EventBus = event.BusDefault
)

type Kernel struct {
	eg         errgroup.Group
	b          func() []boot.Bootloader
	p          func() []process.Processer
	stopSignal chan struct{}
}

func NewKernel(b func() []boot.Bootloader, p func() []process.Processer) *Kernel {
	return &Kernel{
		eg:         errgroup.Group{},
		p:          p,
		b:          b,
		stopSignal: make(chan struct{}),
	}
}

func (a *Kernel) Start(ctx context.Context) {
	go func() {
		select {
		case <-signalc.Wait(syscall.SIGINT, syscall.SIGTERM):
			close(a.stopSignal)
		case <-ctx.Done():
			close(a.stopSignal)
		}
	}()

	boot.MustBoot(a.b())

	for _, processer := range a.p() {
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
