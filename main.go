package main

import (
	"github.com/kingparks/cursor-vip/auth"
	"github.com/kingparks/cursor-vip/tui"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	productIndexSelected, modelIndexSelected := tui.Run()
	startServer(productIndexSelected, modelIndexSelected)
}

func startServer(productIndexSelected string, modelIndexSelected int) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		tui.UnSetProxy()
		if modelIndexSelected == 1 {
			auth.UnSetClient(productIndexSelected)
		}
		os.Exit(0)
	}()
	if modelIndexSelected == 1 {
		tui.SetProxy("localhost", auth.Port)
	}
	auth.Run(productIndexSelected)
}
