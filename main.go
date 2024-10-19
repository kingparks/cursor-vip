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
	tui.Sigs = make(chan os.Signal, 1)
	signal.Notify(tui.Sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-tui.Sigs
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
