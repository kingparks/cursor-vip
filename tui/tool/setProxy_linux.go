//go:build linux
// +build linux

package tool

import "fmt"
import "os"
import "syscall"

func SetProxy(server string, port string) {
	fmt.Println("Set HTTP and HTTPS proxy manually: ", server+":"+port)
}

func UnSetProxy() {
	fmt.Println("UnSet HTTP and HTTPS proxy manually: ")
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
	return true
}
