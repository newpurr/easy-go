package event

import (
	"context"
	evbus "github.com/asaskevich/EventBus"
)

var (
	BusDefault Bus = evbus.New()
)

type (
	EventHandler func(ctx EventContext)
	EventContext struct {
		context.Context
		Publisher interface{}
		Param     interface{}
	}
	Subscriber struct {
		Topic string
		Fn    EventHandler
	}
)

//BusSubscriber defines subscription-related bus behavior
type BusSubscriber interface {
	Subscribe(topic string, fn interface{}) error
	SubscribeAsync(topic string, fn interface{}, transactional bool) error
	Unsubscribe(topic string, handler interface{}) error
}

//BusPublisher defines publishing-related bus behavior
type BusPublisher interface {
	Publish(topic string, args ...interface{})
}

//Bus englobes global (subscribe, publish) bus behavior
type Bus interface {
	BusSubscriber
	BusPublisher
}
