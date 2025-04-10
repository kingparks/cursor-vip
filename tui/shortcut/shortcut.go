package shortcut

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
	"github.com/kingparks/cursor-vip/tui/client"
	"github.com/kingparks/cursor-vip/tui/params"
	"github.com/kingparks/cursor-vip/tui/tool"
	"strings"
	"syscall"
)

var payUrl string
var orderIDExclusive string
var orderIDU3d string
var orderIDU3t string
var orderIDU3h string

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
			//fmt.Println("Error reading keyboard:", err)
			return
		}

		//// 检查是否按下 Ctrl+C
		if key == keyboard.KeyCtrlC {
			// 发送退出信号
			params.Sigs <- syscall.SIGTERM
			keyboard.Close()
		}
		// 判断是否按下回车键
		if key == keyboard.KeyEnter {
			//break
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
			params.Lang = "en"
			tool.SetConfig(params.Lang, params.Mode)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("Settings successful, will take effect after restart"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "szh"):
			params.Lang = "zh"
			tool.SetConfig(params.Lang, params.Mode)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("Settings successful, will take effect after restart"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "sm1"):
			params.Mode = 1
			tool.SetConfig(params.Lang, params.Mode)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("设置成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "sm2"):
			params.Mode = 2
			tool.SetConfig(params.Lang, params.Mode)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("设置成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "sm3"):
			params.Mode = 3
			tool.SetConfig(params.Lang, params.Mode)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("设置成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "sm4"):
			params.Mode = 4
			tool.SetConfig(params.Lang, params.Mode)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("设置成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "buy"):
			payUrl, orderIDExclusive = client.Cli.GetExclusivePayUrl()
			_ = clipboard.WriteAll(payUrl)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.DGreen, payUrl)
			fmt.Println(params.Trr.Tr("捐赠完成后请依次按键 ckp"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "u3d"):
			if !client.Cli.CheckFToken(params.DeviceID) {
				_, _ = fmt.Fprintf(params.ColorOut, params.Yellow, params.Trr.Tr("抱歉，模式三暂无新账号，请稍后再试"))
				return
			}
			if params.M3c > "0" {
				if err = client.Cli.DelFToken(params.DeviceID, "u3"); err != nil {
					return
				}
				_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
				tool.OpenNewTerminal()
				return
			}
			payUrl, orderIDU3d = client.Cli.GetM3PayUrl()
			_ = clipboard.WriteAll(payUrl)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.DGreen, payUrl)
			fmt.Println(params.Trr.Tr("捐赠完成后请依次按键 c3p"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "u3t"):
			if !client.Cli.CheckFToken(params.DeviceID) {
				_, _ = fmt.Fprintf(params.ColorOut, params.Yellow, params.Trr.Tr("抱歉，模式三暂无新账号，请稍后再试"))
				return
			}
			if params.M3c > "0" {
				if err = client.Cli.DelFToken(params.DeviceID, "u3"); err != nil {
					return
				}
				_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
				tool.OpenNewTerminal()
				return
			}
			payUrl, orderIDU3t = client.Cli.GetM3tPayUrl()
			_ = clipboard.WriteAll(payUrl)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.DGreen, payUrl)
			fmt.Println(params.Trr.Tr("捐赠完成后请依次按键 c3t"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "u3h"):
			if !client.Cli.CheckFToken(params.DeviceID) {
				_, _ = fmt.Fprintf(params.ColorOut, params.Yellow, params.Trr.Tr("抱歉，模式三暂无新账号，请稍后再试"))
				return
			}
			if params.M3c > "0" {
				if err = client.Cli.DelFToken(params.DeviceID, "u3"); err != nil {
					return
				}
				_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
				tool.OpenNewTerminal()
				return
			}
			payUrl, orderIDU3h = client.Cli.GetM3hPayUrl()
			_ = clipboard.WriteAll(payUrl)
			fmt.Println()
			_, _ = fmt.Fprintf(params.ColorOut, params.DGreen, payUrl)
			fmt.Println(params.Trr.Tr("捐赠完成后请依次按键 c3h"))
			keyBuffer = nil

		case strings.HasSuffix(combination, "ckp"):
			fmt.Println("checking...")
			isPay := client.Cli.ExclusivePayCheck(orderIDExclusive, params.DeviceID)
			if !isPay {
				fmt.Println(params.Trr.Tr("未捐赠,请捐赠完成后回车"))
				continue
			}
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "c3p"):
			fmt.Println("checking...")
			isPay := client.Cli.M3PayCheck(orderIDU3d, params.DeviceID)
			if !isPay {
				fmt.Println(params.Trr.Tr("未捐赠,请捐赠完成后回车"))
				continue
			}
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "c3t"):
			fmt.Println("checking...")
			isPay := client.Cli.M3tPayCheck(orderIDU3t, params.DeviceID)
			if !isPay {
				fmt.Println(params.Trr.Tr("未捐赠,请捐赠完成后回车"))
				continue
			}
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "c3h"):
			fmt.Println("checking...")
			isPay := client.Cli.M3hPayCheck(orderIDU3h, params.DeviceID)
			if !isPay {
				fmt.Println(params.Trr.Tr("未捐赠,请捐赠完成后回车"))
				continue
			}
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("购买成功，将在重启 cursor-vip 后生效"))
			keyBuffer = nil
			tool.OpenNewTerminal()

		case strings.HasSuffix(combination, "q3d"):
			_, _ = fmt.Fprintf(params.ColorOut, params.Green, "\t"+params.M3DaysRemainingOnTrial+"d")
			keyBuffer = nil
		}
	}
}
