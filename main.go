package main

import (
	"github.com/kingparks/cursor-vip/auth"
	"github.com/kingparks/cursor-vip/tui"
	"github.com/kingparks/cursor-vip/tui/params"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	productSelected, modelIndexSelected := tui.Run()
	startServer(productSelected, modelIndexSelected)
}

func startServer(productSelected string, modelIndexSelected int) {
	params.Sigs = make(chan os.Signal, 1)
	signal.Notify(params.Sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-params.Sigs
		auth.UnSetClient(productSelected)
		os.Exit(0)
	}()
	auth.Run(productSelected, modelIndexSelected)
}
