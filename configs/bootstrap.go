package configs

import (
	"github.com/newpurr/easy-go/pkg/boot"
)

var (
	AppBootstraps = []boot.Bootloader{
		// The launcher the application must rely on, please do not edit
		boot.NewSettingBootloader("configs/" + Env),
		boot.NewEventBusBootloader(Events),
		boot.NewLoggerBootloader(),
		boot.NewValidatorBootloader(),

		// optional launcher
		boot.NewGinBootloader(),
		boot.NewJaegerTracerBootloader(),
		boot.NewMysqlBootloader(),

		// add your custom launcher
		// ...
	}
)
