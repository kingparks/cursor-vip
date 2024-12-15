package shortcut

import (
	"fmt"
	"github.com/kingparks/cursor-vip/tui/params"
	"github.com/kingparks/cursor-vip/tui/tool"
	hook "github.com/robotn/gohook"
)

func Do() {
	// 设置语言为英文
	hook.Register(hook.KeyDown, []string{"s", "e", "n"}, func(e hook.Event) {
		tool.SetConfig("en", params.Mode)
		fmt.Printf(params.Red, params.Trr.Tr("Settings successful, will take effect after manual restart"))
	})
	// 设置语言为中文
	hook.Register(hook.KeyDown, []string{"s", "z", "h"}, func(e hook.Event) {
		tool.SetConfig("zh", params.Mode)
		fmt.Printf(params.Red, params.Trr.Tr("Settings successful, will take effect after manual restart"))
	})
	// 设置为普通模式
	hook.Register(hook.KeyDown, []string{"s", "m", "1"}, func(e hook.Event) {
		tool.SetConfig(params.Lang, 1)
		fmt.Printf(params.Red, params.Trr.Tr("设置成功，将在手动重启后生效"))
	})
	// 设置为代理模式
	hook.Register(hook.KeyDown, []string{"s", "m", "2"}, func(e hook.Event) {
		tool.SetConfig(params.Lang, 2)
		fmt.Printf(params.Red, params.Trr.Tr("设置成功，将在手动重启后生效"))
	})
	// 普通
	s := hook.Start()
	<-hook.Process(s)
}
