package configs

import (
	"github.com/newpurr/easy-go/pkg/app"
	"github.com/newpurr/easy-go/pkg/boot"
	"log"
)

var (
	Events = []boot.Event{
		{
			app.HttpServerBeforeStartTopic,
			func(eventContext boot.EventContext) {
				log.Println("HttpServer beforeStart event handler")
			},
		},
		{
			app.HttpServerBeforeStopTopic,
			func(eventContext boot.EventContext) {
				log.Println("HttpServer beforeStop event handler")
			},
		},
		{
			app.HttpServerAfterStopTopic,
			func(eventContext boot.EventContext) {
				log.Println("HttpServer afterStop event handler")
			},
		},
	}
)
