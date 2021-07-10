package boot

import (
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/pkg/tracer"
)

type JaegerTracerBootloader struct {
}

func NewJaegerTracerBootloader() *JaegerTracerBootloader {
	return &JaegerTracerBootloader{}
}

func (sb JaegerTracerBootloader) Boot() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	application.Tracer = jaegerTracer
	return nil
}
