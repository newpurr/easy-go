package boot

import (
	"github.com/newpurr/easy-go/event"
)

type EventSubscribeBootloader struct {
	b      event.Bus
	Events []event.Subscriber
}

func NewEventSubscribeBootloader(b event.Bus, events []event.Subscriber) *EventSubscribeBootloader {
	return &EventSubscribeBootloader{Events: events, b: b}
}

func (b *EventSubscribeBootloader) Boot() error {
	for _, e := range b.Events {
		_ = b.b.Subscribe(e.Topic, e.Fn)
	}

	return nil
}
