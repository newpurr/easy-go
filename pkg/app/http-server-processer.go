package app

import (
	"context"
	"github.com/newpurr/easy-go/pkg/boot"
	"github.com/newpurr/easy-go/pkg/event"
	"log"
	"net/http"
	"time"
)

var (
	HttpServerBeforeStartTopic = "HttpServer.beforeStart"
	HttpServerErrorStartTopic  = "HttpServer.startError"
	HttpServerBeforeStopTopic  = "HttpServer.beforeStop"
	HttpServerAfterStopTopic   = "HttpServer.afterStop"
)

type HttpServer struct {
	*http.Server
	EventBus event.Bus
}

func NewHttpServer(server *http.Server, eventBus event.Bus) *HttpServer {
	if server == nil {
		log.Fatalln("server cannot be nil")
	}
	if eventBus == nil {
		log.Fatalln("eventBus cannot be nil")
	}

	return &HttpServer{Server: server, EventBus: eventBus}
}

func (s HttpServer) MustStart() {
	s.EventBus.Publish(HttpServerBeforeStartTopic, boot.EventContext{})
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.EventBus.Publish(HttpServerErrorStartTopic, boot.EventContext{})
		log.Fatalf("s.ListenAndServe err: %v", err)
	}
}

func (s HttpServer) Stop() {
	s.EventBus.Publish(HttpServerBeforeStopTopic, boot.EventContext{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer s.EventBus.Publish(HttpServerAfterStopTopic, boot.EventContext{})
	err := s.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Shuting down server...")
}
