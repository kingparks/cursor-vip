//go:build darwin
// +build darwin

package tool

import (
	"log"
	"os"
	"os/exec"
	"syscall"
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

// CheckIfChownRequired 检查是否需要执行 chown 操作
func CheckIfChownRequired(filePath string) bool {
	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	// 获取文件的UID
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}
	// 获取当前用户的UID
	currentUID := os.Getuid()
	// 比较文件UID和当前用户UID
	if int(stat.Uid) != currentUID {
		return true
	}
	return false
}

func MacOSIsSIPDisable() bool {
	// 执行 csrutil status 命令，检查 SIP 状态
	// 如果 SIP 状态为已禁用，则返回 true
	// 否则返回 false
	res, err := exec.Command("csrutil", "status").Output()
	if err != nil {
		return false
	}
	if string(res) == "System Integrity Protection status: disabled.\n" {
		return true
	}
	return false
}
