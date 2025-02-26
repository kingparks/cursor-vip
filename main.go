package main

import (
	"github.com/gofrs/flock"
	"github.com/kingparks/cursor-vip/auth"
	"github.com/kingparks/cursor-vip/tui"
	"github.com/kingparks/cursor-vip/tui/params"
	"github.com/kingparks/cursor-vip/tui/shortcut"
	"github.com/kingparks/cursor-vip/tui/tool"
	"os"
	"os/signal"
	"syscall"
)

var lock *flock.Flock
var pidFilePath string

func main() {
	lock, pidFilePath, _ = tool.EnsureSingleInstance("cursor-vip")
	productSelected, modelIndexSelected := tui.Run()
	startServer(productSelected, modelIndexSelected)
}

func startServer(productSelected string, modelIndexSelected int) {
	params.Sigs = make(chan os.Signal, 1)
	signal.Notify(params.Sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-params.Sigs
		_ = lock.Unlock()
		_ = os.Remove(pidFilePath)
		auth.UnSetClient(productSelected)
		os.Exit(0)
	}()
	go shortcut.Do()
	auth.Run(productSelected, modelIndexSelected)
}
