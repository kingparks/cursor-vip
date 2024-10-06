package main

import (
	"github.com/kingparks/cursor-vip/auth"
	"github.com/kingparks/cursor-vip/tui"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	productSelected, modelIndexSelected := tui.Run()
	startServer(productSelected, modelIndexSelected)
}

func startServer(productSelected string, modelIndexSelected int) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-sigs
		auth.UnSetClient(productSelected)
		if modelIndexSelected == 2 {
			tui.UnSetProxy()
		}
		os.Exit(0)
	}()
	if modelIndexSelected == 2 {
		tui.SetProxy("localhost", auth.Port)
	}
	auth.Run(productSelected, modelIndexSelected)
}
