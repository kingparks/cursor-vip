package tui

import (
	"embed"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/kingparks/cursor-vip/tui/client"
	"github.com/kingparks/cursor-vip/tui/params"
	"github.com/kingparks/cursor-vip/tui/shortcut"
	"github.com/kingparks/cursor-vip/tui/tool"
	"github.com/mattn/go-colorable"

	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/unknwon/i18n"
)

//go:embed all:locales
var localeFS embed.FS

// Run 启动
func Run() (productSelected string, modelIndexSelected int) {
	params.ColorOut = colorable.NewColorableStdout()
	params.Lang, params.Promotion, params.Mode = tool.GetConfig()
	params.DeviceID = tool.GetMachineID()
	params.MachineID = tool.GetMachineID()
	client.Cli = client.Client{Hosts: params.Hosts}

	localeFileEn, _ := localeFS.ReadFile("locales/en.ini")
	_ = i18n.SetMessage("en", localeFileEn)
	localeFileNl, _ := localeFS.ReadFile("locales/nl.ini")
	_ = i18n.SetMessage("nl", localeFileNl)
	localeFileRu, _ := localeFS.ReadFile("locales/ru.ini")
	_ = i18n.SetMessage("ru", localeFileRu)
	localeFileHu, _ := localeFS.ReadFile("locales/hu.ini")
	_ = i18n.SetMessage("hu", localeFileHu)
	localeFileTr, _ := localeFS.ReadFile("locales/tr.ini")
	_ = i18n.SetMessage("tr", localeFileTr)
	localeFileEs, _ := localeFS.ReadFile("locales/es.ini")
	_ = i18n.SetMessage("es", localeFileEs)
	switch params.Lang {
	case "zh":
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "zh"}}
		params.GithubPath = "https://ghp.ci/" + params.GithubPath
	case "nl":
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "nl"}}
	case "ru":
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "ru"}}
	case "hu":
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "hu"}}
	case "tr":
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "tr"}}
	case "es":
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "es"}}
	default:
		params.Trr = &params.Tr{Locale: i18n.Locale{Lang: "en"}}
	}

	_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("CURSOR VIP")+` v`+strings.Join(strings.Split(fmt.Sprint(params.Version), ""), "."))
	// 检查是否在容器环境
	if content, err := os.ReadFile("/proc/1/cgroup"); err == nil {
		if strings.Contains(string(content), "/docker/") {
			_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("不支持容器环境"))
			_, _ = fmt.Scanln()
			panic(params.Trr.Tr("不支持容器环境"))
		}
	}
	client.Cli.SetProxy(params.Lang)
	checkUpdate(params.Version)
	sCount, sPayCount, _, _, exp := client.Cli.GetMyInfo(params.DeviceID)
	_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("设备码")+":"+params.DeviceID)
	expTime, _ := time.ParseInLocation("2006-01-02 15:04:05", exp, time.Local)
	_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("付费到期时间")+":"+exp)
	_, _ = fmt.Fprintf(params.ColorOut, "\033[32m%s\033[0m\u001B[1;32m %s \u001B[0m\033[32m%s\033[0m\u001B[1;32m %s \u001B[0m\u001B[32m%s\u001B[0m\n",
		params.Trr.Tr("推广命令：(已推广"), sCount, params.Trr.Tr("人,推广已付费"), sPayCount, params.Trr.Tr("人；每推广10人或推广付费2人可获得一年授权)"))
	_, _ = fmt.Fprintf(params.ColorOut, params.HGreen, "bash <(curl -Lk "+params.GithubPath+"install.sh) "+params.DeviceID+"\n")
	_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("专属推广链接")+"："+params.Host+"?p="+params.DeviceID)
	fmt.Println()

	printAD()
	fmt.Println()

	// 快捷键
	_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("Switch to English：simultaneously press keyboard 's' 'e' 'n'"))
	modelIndexSelected = int(params.Mode)
	switch params.Mode {
	case 1:
		_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("切换为代理模式：同时按键盘 's' 'm' '2'"))
	case 2:
		_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("切换为普通模式：同时按键盘 's' 'm' '1'"))
	}
	go shortcut.Do()
	fmt.Println()

	if len(params.Product) > 1 {
		_, _ = fmt.Fprintf(params.ColorOut, params.DefaultColor, params.Trr.Tr("选择要授权的产品："))
		for i, v := range params.Product {
			_, _ = fmt.Fprintf(params.ColorOut, params.HGreen, fmt.Sprintf("%d. %s\t", i+1, v))
		}
		fmt.Println()
		fmt.Print(params.Trr.Tr("请输入产品编号（直接回车默认为1，可以同时输入多个例如 145）："))
		productIndex := 1
		_, _ = fmt.Scanln(&productIndex)
		if productIndex < 1 {
			fmt.Println(params.Trr.Tr("输入有误"))
			return
		}
		for _, v := range strings.Split(fmt.Sprint(productIndex), "") {
			vi, _ := strconv.Atoi(v)
			productSelected += params.Product[vi-1] + ","
		}
		if len(productSelected) > 1 {
			productSelected = productSelected[:len(productSelected)-1]
		}
		fmt.Println(params.Trr.Tr("选择的产品为：") + productSelected)
		fmt.Println()
	} else {
		productSelected = params.Product[0]
	}
	// 到期了
	periodIndex := 1
	if expTime.Before(time.Now()) {
		_, _ = fmt.Fprintf(params.ColorOut, params.DefaultColor, params.Trr.Tr("选择有效期："))
		//jbPeriod := []string{"1" + params.Trr.Tr("年(购买)"), "2" + params.Trr.Tr("小时(免费)")}
		jbPeriod := []string{"1" + params.Trr.Tr("年(购买)")}
		for i, v := range jbPeriod {
			_, _ = fmt.Fprintf(params.ColorOut, params.HGreen, fmt.Sprintf("%d. %s\t", i+1, v))
		}
		fmt.Println()
		_, _ = fmt.Fprintf(params.ColorOut, "%s", params.Trr.Tr("请输入有效期编号（直接回车默认为1）："))
		_, _ = fmt.Scanln(&periodIndex)
		if periodIndex < 1 || periodIndex > len(jbPeriod) {
			fmt.Println(params.Trr.Tr("输入有误"))
			return
		}
		fmt.Println(params.Trr.Tr("选择的有效期为：") + jbPeriod[periodIndex-1])
		fmt.Println()

		//if periodIndex == 2 {
		//	_, _ = fmt.Fprintf(params.ColorOut, green, Trr.Tr("授权成功！使用过程请不要关闭此窗口"))
		//	countDown(2 * 60 * 60)
		//	return
		//}

		payUrl, orderID := client.Cli.GetPayUrl()
		isCopyText := ""
		errClip := clipboard.WriteAll(payUrl)
		if errClip == nil {
			isCopyText = params.Trr.Tr("（已复制到剪贴板）")
		}
		fmt.Println(params.Trr.Tr("付费已到期,捐赠以获取一年期授权") + isCopyText)
		_, _ = fmt.Fprintf(params.ColorOut, params.DGreen, payUrl)
		fmt.Println(params.Trr.Tr("捐赠完成后请回车"))
		//检测控制台回车
	checkPay:
		_, _ = fmt.Scanln()
		isPay := client.Cli.PayCheck(orderID, params.DeviceID)
		if !isPay {
			fmt.Println(params.Trr.Tr("未捐赠,请捐赠完成后回车"))
			goto checkPay
		}
		_, _, _, _, exp = client.Cli.GetMyInfo(params.DeviceID)
		expTime, _ = time.ParseInLocation("2006-01-02 15:04:05", exp, time.Local)
		fmt.Println()
	}
	_, _ = fmt.Fprintf(params.ColorOut, params.Green, params.Trr.Tr("授权成功！使用过程请不要关闭此窗口"))
	tool.CountDown(int(expTime.Sub(time.Now()).Seconds()))
	return
}

func printAD() {
	ad := client.Cli.GetAD()
	if len(ad) == 0 {
		return
	}
	_, _ = fmt.Fprintf(params.ColorOut, params.Yellow, ad)
}

func checkUpdate(version int) {
	upUrl := client.Cli.CheckVersion(fmt.Sprint(version))
	if upUrl == "" {
		return
	}
	isCopyText := ""
	installCmd := `bash -c "$(curl -fsSLk ` + params.GithubPath + `install.sh)"`
	errClip := clipboard.WriteAll(installCmd)
	if errClip == nil {
		isCopyText = params.Trr.Tr("（已复制到剪贴板）")
	}
	switch runtime.GOOS {
	case "windows":
		_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("有新版本，请关闭本窗口，将下面命令粘贴到GitBash窗口执行")+isCopyText+`：`)
	default:
		_, _ = fmt.Fprintf(params.ColorOut, params.Red, params.Trr.Tr("有新版本，请关闭本窗口，将下面命令粘贴到新终端窗口执行")+isCopyText+`：`)
	}
	_, _ = fmt.Fprintf(params.ColorOut, params.HGreen, installCmd)
	_, _ = fmt.Scanln()
	os.Exit(0)
	return
}
