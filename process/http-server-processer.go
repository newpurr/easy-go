package process

import (
	"context"
	"github.com/newpurr/easy-go/event"
	"github.com/newpurr/easy-go/httpc"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	*http.Server
	EventBus event.Bus
}

func NewHttpServerProcess(server *http.Server, eventBus event.Bus) *HttpServer {
	if server == nil {
		log.Fatalln("server cannot be nil")
	}
	if eventBus == nil {
		log.Fatalln("eventBus cannot be nil")
	}

	return &HttpServer{Server: server, EventBus: eventBus}
}

func (s *HttpServer) MustStart() {
	s.EventBus.Publish(httpc.ServerBeforeStartTopic, s.eventContext())
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.EventBus.Publish(httpc.ServerErrorStartTopic, s)
		log.Fatalf("s.ListenAndServe err: %v", err)
	}
}

func (s *HttpServer) Stop() {
	s.EventBus.Publish(httpc.ServerBeforeStopTopic, s.eventContext())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer s.EventBus.Publish(httpc.ServerAfterStopTopic, s.eventContext())
	err := s.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Shuting down server...")
}

func (s *HttpServer) eventContext() event.EventContext {
	return event.EventContext{
		Publisher: s,
	}
}
