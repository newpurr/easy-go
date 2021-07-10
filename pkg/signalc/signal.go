package signalc

import (
	"os"
	"os/signal"
)

func Wait(sig ...os.Signal) chan os.Signal {
	quit := make(chan os.Signal)
	signal.Notify(quit, sig...)
	return quit
}
