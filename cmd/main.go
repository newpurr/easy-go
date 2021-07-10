package main

import (
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/configs"
	"github.com/newpurr/easy-go/internal/routers"
	"github.com/newpurr/easy-go/pkg/app"
	"github.com/newpurr/easy-go/pkg/assert"
	"github.com/newpurr/easy-go/pkg/boot"
	_ "go.uber.org/automaxprocs"
	"net/http"
)

func init() {
	boot.MustBoot(configs.AppBootstraps)
	assert.MustTrue(application.ServerSetting != nil, "the service configuration must be loaded first")
}

func main() {
	httpServerProcess := app.NewHttpServer(&http.Server{
		Addr:           ":" + application.ServerSetting.HttpPort,
		Handler:        routers.NewRouter(),
		ReadTimeout:    application.ServerSetting.ReadTimeout,
		WriteTimeout:   application.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}, application.EventBus)

	antsGoroutineProcess := app.NewAntsGoroutineProcess(5, 10)

	NewAntsMnsGoroutineProcess := app.NewAntsMnsGoroutineProcess()

	app.FastStart(httpServerProcess, antsGoroutineProcess, NewAntsMnsGoroutineProcess)
}
