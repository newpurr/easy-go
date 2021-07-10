package boot

import (
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/pkg/event"
)

type (
	EventHandler func(EventContext)
	EventContext struct {
		Param interface{}
	}
	Event struct {
		Topic string
		Fn    EventHandler
	}
)

type EventBusBootloader struct {
	Events []Event
}

func NewEventBusBootloader(events []Event) *EventBusBootloader {
	return &EventBusBootloader{Events: events}
}

func (b EventBusBootloader) Boot() error {
	application.EventBus = event.BusDefault

	for _, e := range b.Events {
		_ = application.EventBus.Subscribe(e.Topic, e.Fn)
	}

	return nil
}
