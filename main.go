package main

import (
	"github.com/kingparks/cursor-vip/auth"
	"github.com/kingparks/cursor-vip/tui"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	productIndexSelected := tui.Run()
	startServer(productIndexSelected)
}

func startServer(productIndexSelected string) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		tui.UnSetProxy()
		os.Exit(0)
	}()
	tui.SetProxy("localhost", auth.Port)
	auth.Run(productIndexSelected)
}
