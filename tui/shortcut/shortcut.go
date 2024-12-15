package shortcut

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/kingparks/cursor-vip/tui/params"
	"github.com/kingparks/cursor-vip/tui/tool"
)

func Do() {
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard:", err)
		return
	}
	defer keyboard.Close()

	var keyBuffer []rune
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading keyboard:", err)
			return
		}

		// 检查是否按下 Ctrl+C
		if key == keyboard.KeyCtrlC {
			return
		}

		// 将按键添加到缓冲区
		if char != 0 {
			keyBuffer = append(keyBuffer, char)
		}

		// 保持缓冲区最多3个字符
		if len(keyBuffer) > 3 {
			keyBuffer = keyBuffer[1:]
		}

		// 检查快捷键组合
		combination := string(keyBuffer)

		switch {
		case strings.HasSuffix(combination, "sen"):
			tool.SetConfig("en", params.Mode)
			fmt.Println()
			fmt.Printf(params.Red, params.Trr.Tr("Settings successful, will take effect after manual restart"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "szh"):
			tool.SetConfig("zh", params.Mode)
			fmt.Println()
			fmt.Printf(params.Red, params.Trr.Tr("Settings successful, will take effect after manual restart"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "sm1"):
			tool.SetConfig(params.Lang, 1)
			fmt.Println()
			fmt.Printf(params.Red, params.Trr.Tr("设置成功，将在手动重启后生效"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "sm2"):
			tool.SetConfig(params.Lang, 2)
			fmt.Println()
			fmt.Printf(params.Red, params.Trr.Tr("设置成功，将在手动重启后生效"))
			keyBuffer = nil

		}
	}
}
