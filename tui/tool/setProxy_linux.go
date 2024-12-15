//go:build linux
// +build linux

package tool

import "fmt"

func SetProxy(server string, port string) {
	fmt.Println("Set HTTP and HTTPS proxy manually: ", server+":"+port)
}

func UnSetProxy() {
	fmt.Println("UnSet HTTP and HTTPS proxy manually: ")
}
