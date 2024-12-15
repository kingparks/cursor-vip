//go:build darwin
// +build darwin

package tool

import (
	"log"
	"os/exec"
)

func SetProxy(server string, port string) {
	// Set HTTP proxy
	cmd := exec.Command("networksetup", "-setwebproxy", "Wi-Fi", server, port)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to set HTTP proxy: %v", err)
	}

	// Set HTTPS proxy
	cmd = exec.Command("networksetup", "-setsecurewebproxy", "Wi-Fi", server, port)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to set HTTPS proxy: %v", err)
	}
}

func UnSetProxy() {
	// Unset HTTP proxy
	cmd := exec.Command("networksetup", "-setwebproxystate", "Wi-Fi", "off")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to unset HTTP proxy: %v", err)
	}

	// Unset HTTPS proxy
	cmd = exec.Command("networksetup", "-setsecurewebproxystate", "Wi-Fi", "off")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to unset HTTPS proxy: %v", err)
	}
}
