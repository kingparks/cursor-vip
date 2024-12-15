//go:build windows
// +build windows

package tool

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

func SetProxy(server string, port string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	err = k.SetStringValue("ProxyServer", server+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	err = k.SetDWordValue("ProxyEnable", 1)
	if err != nil {
		log.Fatal(err)
	}
}

func UnSetProxy() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	err = k.SetDWordValue("ProxyEnable", 0)
	if err != nil {
		log.Fatal(err)
	}
}
