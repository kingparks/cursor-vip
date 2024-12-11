package tui

import (
	"crypto/md5"
	"embed"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/denisbrodbeck/machineid"
	"howett.net/plist"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/unknwon/i18n"
)

var version = 229

var hosts = []string{"https://cursor.jeter.eu.org", "http://129.154.205.7:7193"}
var host = hosts[0]
var githubPath = "https://ghp.ci/https://github.com/kingparks/cursor-vip/releases/download/latest/"
var err error

var green = "\033[32m%s\033[0m\n"
var yellow = "\033[33m%s\033[0m\n"
var hGreen = "\033[1;32m%s\033[0m"
var dGreen = "\033[4;32m%s\033[0m\n"
var red = "\033[31m%s\033[0m\n"
var defaultColor = "%s"
var lang, _ = getLocale()
var deviceID = getMacMD5_241018()
var machineID = getMacMD5_241019()
var Cli = Client{Hosts: hosts}
var Sigs chan os.Signal

//go:embed all:locales
var localeFS embed.FS

type Tr struct {
	i18n.Locale
}

var Trr *Tr

var jbProduct = []string{"cursor IDE"}

func Run() (productSelected string, modelIndexSelected int) {
	language := flag.String("l", lang, "set language, eg: zh, en, nl, ru, hu, Trr")
	flag.Parse()

	localeFileEn, _ := localeFS.ReadFile("locales/en.ini")
	_ = i18n.SetMessage("en", localeFileEn)
	localeFileNl, _ := localeFS.ReadFile("locales/nl.ini")
	_ = i18n.SetMessage("nl", localeFileNl)
	localeFileRu, _ := localeFS.ReadFile("locales/ru.ini")
	_ = i18n.SetMessage("ru", localeFileRu)
	localeFileHu, _ := localeFS.ReadFile("locales/hu.ini")
	_ = i18n.SetMessage("hu", localeFileHu)
	localeFileTr, _ := localeFS.ReadFile("locales/Trr.ini")
	_ = i18n.SetMessage("Trr", localeFileTr)
	localeFileEs, _ := localeFS.ReadFile("locales/es.ini")
	_ = i18n.SetMessage("es", localeFileEs)
	lang = *language
	switch lang {
	case "zh":
		Trr = &Tr{Locale: i18n.Locale{Lang: "zh"}}
	case "nl":
		Trr = &Tr{Locale: i18n.Locale{Lang: "nl"}}
	case "ru":
		Trr = &Tr{Locale: i18n.Locale{Lang: "ru"}}
	case "hu":
		Trr = &Tr{Locale: i18n.Locale{Lang: "hu"}}
	case "Trr":
		Trr = &Tr{Locale: i18n.Locale{Lang: "Trr"}}
	case "es":
		Trr = &Tr{Locale: i18n.Locale{Lang: "es"}}
	default:
		Trr = &Tr{Locale: i18n.Locale{Lang: "en"}}
	}

	fmt.Printf(green, Trr.Tr("CURSOR VIP")+` v`+strings.Join(strings.Split(fmt.Sprint(version), ""), "."))
	// 检查是否在容器环境
	if content, err := os.ReadFile("/proc/1/cgroup"); err == nil {
		if strings.Contains(string(content), "/docker/") {
			fmt.Printf(red, Trr.Tr("不支持容器环境"))
			_, _ = fmt.Scanln()
			panic(Trr.Tr("不支持容器环境"))
		}
	}
	Cli.SetProxy(lang)
	checkUpdate(version)
	sCount, sPayCount, _, _, exp := Cli.GetMyInfo(deviceID)
	fmt.Printf(green, Trr.Tr("设备码")+":"+deviceID)
	expTime, _ := time.ParseInLocation("2006-01-02 15:04:05", exp, time.Local)
	fmt.Printf(green, Trr.Tr("付费到期时间")+":"+exp)
	fmt.Printf("\033[32m%s\033[0m\u001B[1;32m %s \u001B[0m\033[32m%s\033[0m\u001B[1;32m %s \u001B[0m\u001B[32m%s\u001B[0m\n",
		Trr.Tr("推广命令：(已推广"), sCount, Trr.Tr("人,推广已付费"), sPayCount, Trr.Tr("人；每推广10人或推广付费2人可获得一年授权)"))
	fmt.Printf(hGreen, "bash <(curl -Lk "+githubPath+"install.sh) "+deviceID+"\n")
	fmt.Printf(green, Trr.Tr("专属推广链接")+"："+host+"?p="+deviceID)
	fmt.Println()

	printAD()
	fmt.Println()

	if false {
		fmt.Printf(defaultColor, Trr.Tr("选择启动模式："))
		for i, v := range []string{Trr.Tr("极简模式"), Trr.Tr("强劲代理模式")} {
			fmt.Printf(hGreen, fmt.Sprintf("%d. %s\t", i+1, v))
		}
		fmt.Println()
		fmt.Print(Trr.Tr("请输入模式编号（直接回车默认为1）："))
		modelIndexSelected = 1
		_, _ = fmt.Scanln(&modelIndexSelected)
		if modelIndexSelected < 1 || modelIndexSelected > 2 {
			fmt.Println(Trr.Tr("输入有误"))
			return
		}
		fmt.Println()
	} else {
		modelIndexSelected = 1
	}

	if len(jbProduct) > 1 {
		fmt.Printf(defaultColor, Trr.Tr("选择要授权的产品："))
		for i, v := range jbProduct {
			fmt.Printf(hGreen, fmt.Sprintf("%d. %s\t", i+1, v))
		}
		fmt.Println()
		fmt.Print(Trr.Tr("请输入产品编号（直接回车默认为1，可以同时输入多个例如 145）："))
		productIndex := 1
		_, _ = fmt.Scanln(&productIndex)
		if productIndex < 1 {
			fmt.Println(Trr.Tr("输入有误"))
			return
		}
		for _, v := range strings.Split(fmt.Sprint(productIndex), "") {
			vi, _ := strconv.Atoi(v)
			productSelected += jbProduct[vi-1] + ","
		}
		if len(productSelected) > 1 {
			productSelected = productSelected[:len(productSelected)-1]
		}
		fmt.Println(Trr.Tr("选择的产品为：") + productSelected)
		fmt.Println()
	} else {
		productSelected = jbProduct[0]
	}
	// 到期了
	periodIndex := 1
	if expTime.Before(time.Now()) {
		fmt.Printf(defaultColor, Trr.Tr("选择有效期："))
		jbPeriod := []string{"1" + Trr.Tr("年(购买)"), "2" + Trr.Tr("小时(免费)")}
		for i, v := range jbPeriod {
			fmt.Printf(hGreen, fmt.Sprintf("%d. %s\t", i+1, v))
		}
		fmt.Println()
		fmt.Printf("%s", Trr.Tr("请输入有效期编号（直接回车默认为1）："))
		_, _ = fmt.Scanln(&periodIndex)
		if periodIndex < 1 || periodIndex > len(jbPeriod) {
			fmt.Println(Trr.Tr("输入有误"))
			return
		}
		fmt.Println(Trr.Tr("选择的有效期为：") + jbPeriod[periodIndex-1])
		fmt.Println()

		if periodIndex == 2 {
			fmt.Printf(green, Trr.Tr("授权成功！使用过程请不要关闭此窗口"))
			countDown(2 * 60 * 60)
			return
		}

		payUrl, orderID := Cli.GetPayUrl()
		isCopyText := ""
		errClip := clipboard.WriteAll(payUrl)
		if errClip == nil {
			isCopyText = Trr.Tr("（已复制到剪贴板）")
		}
		fmt.Println(Trr.Tr("付费已到期,捐赠以获取一年期授权") + isCopyText)
		fmt.Printf(dGreen, payUrl)
		fmt.Println(Trr.Tr("捐赠完成后请回车"))
		//检测控制台回车
	checkPay:
		_, _ = fmt.Scanln()
		isPay := Cli.PayCheck(orderID, deviceID)
		if !isPay {
			fmt.Println(Trr.Tr("未捐赠,请捐赠完成后回车"))
			goto checkPay
		}
		_, _, _, _, exp = Cli.GetMyInfo(deviceID)
		expTime, _ = time.ParseInLocation("2006-01-02 15:04:05", exp, time.Local)
		fmt.Println()
	}
	fmt.Printf(green, Trr.Tr("授权成功！使用过程请不要关闭此窗口"))
	countDown(int(expTime.Sub(time.Now()).Seconds()))
	return
}
func countDown(seconds int) {
	go func(seconds int) {
		countdown := seconds // Countdown in seconds
		for countdown >= 0 {
			days := countdown / (24 * 3600)
			hours := (countdown % (24 * 3600)) / 3600
			minutes := (countdown % 3600) / 60
			seconds := countdown % 60

			fmt.Printf("\r%dd %dh %dm %ds", days, hours, minutes, seconds)
			time.Sleep(1 * time.Second)
			countdown--
		}
		// 发送退出信号
		Sigs <- syscall.SIGTERM
	}(seconds)
}

func getMacMD5() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	var macAddress []string
	var wifiAddress []string
	var bluetoothAddress []string
	var macErrorStr string
	for _, inter := range interfaces {
		// 排除虚拟网卡
		hardwareAddr := inter.HardwareAddr.String()
		if hardwareAddr == "" {
			//fmt.Println(fmt.Sprintf("log: have not hardwareAddr :%+v",inter))
			continue
		}
		macErrorStr += inter.Name + ":" + hardwareAddr + "\n"
		virtualMacPrefixes := []string{
			"00:05:69", "00:0C:29", "00:1C:14", "00:50:56", // VMware
			"00:15:5D",             // Hyper-V
			"08:00:27", "0A:00:27", // VirtualBox
		}
		isVirtual := false
		for _, prefix := range virtualMacPrefixes {
			if strings.HasPrefix(hardwareAddr, strings.ToLower(prefix)) {
				isVirtual = true
				break
			}
		}
		if isVirtual {
			//fmt.Println(fmt.Sprintf("log: isVirtual :%+v",inter))
			continue
		}
		// 大于en6的排除
		if strings.HasPrefix(inter.Name, "en") {
			numStr := inter.Name[2:]
			num, _ := strconv.Atoi(numStr)
			if num > 6 {
				//fmt.Println(fmt.Sprintf("log: is num>6 :%+v",inter))
				continue
			}
		}
		if strings.HasPrefix(inter.Name, "en") || strings.HasPrefix(inter.Name, "Ethernet") || strings.HasPrefix(inter.Name, "以太网") || strings.HasPrefix(inter.Name, "WLAN") {
			//fmt.Println(fmt.Sprintf("log: add :%+v",inter))
			macAddress = append(macAddress, hardwareAddr)
		} else if strings.HasPrefix(inter.Name, "Wi-Fi") || strings.HasPrefix(inter.Name, "无线网络") {
			wifiAddress = append(wifiAddress, hardwareAddr)
		} else if strings.HasPrefix(inter.Name, "Bluetooth") || strings.HasPrefix(inter.Name, "蓝牙网络连接") {
			bluetoothAddress = append(bluetoothAddress, hardwareAddr)
		} else {
			//fmt.Println(fmt.Sprintf("log: not add :%+v",inter))
		}
	}
	if len(macAddress) == 0 {
		macAddress = append(macAddress, wifiAddress...)
		if len(macAddress) == 0 {
			macAddress = append(macAddress, bluetoothAddress...)
		}
		if len(macAddress) == 0 {
			fmt.Printf(red, "no mac address found,Please contact customer service")
			_, _ = fmt.Scanln()
			return macErrorStr
		}
	}
	sort.Strings(macAddress)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(macAddress, ","))))
}

func getMac_241018() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	var macError []string
	for _, inter := range interfaces {
		hardwareAddr := inter.HardwareAddr.String()
		if hardwareAddr == "" {
			continue
		}
		macError = append(macError, inter.Name+": "+hardwareAddr)
	}
	sort.Strings(macError)
	return strings.Join(macError, "\n")
}
func getMacMD5_241018() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	var macAddress, bluetoothAddress []string
	virtualMacPrefixes := []string{
		"00:05:69", "00:0C:29", "00:1C:14", "00:50:56", // VMware
		"00:15:5D",             // Hyper-V
		"08:00:27", "0A:00:27", // VirtualBox
	}

	for _, inter := range interfaces {
		hardwareAddr := inter.HardwareAddr.String()
		if hardwareAddr == "" {
			continue
		}

		isVirtual := false
		for _, prefix := range virtualMacPrefixes {
			if strings.HasPrefix(hardwareAddr, strings.ToLower(prefix)) {
				isVirtual = true
				break
			}
		}
		if isVirtual {
			continue
		}

		switch {
		case strings.HasPrefix(inter.Name, "en"), strings.HasPrefix(inter.Name, "Ethernet"), strings.HasPrefix(inter.Name, "以太网"):
			macAddress = append(macAddress, hardwareAddr)
		case strings.HasPrefix(inter.Name, "Bluetooth"), strings.HasSuffix(inter.Name, "Bluetooth"), strings.HasPrefix(inter.Name, "蓝牙网络连接"):
			bluetoothAddress = append(bluetoothAddress, hardwareAddr)
		}
	}

	if len(macAddress) == 0 {
		macAddress = append(macAddress, bluetoothAddress...)
		if len(macAddress) == 0 {
			//fmt.Printf(red, "no mac address found,Please contact customer service")
			//_, _ = fmt.Scanln()
			//return macErrorStr
			return getMacMD5_241019()
		}
	}
	sort.Strings(macAddress)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(macAddress, ","))))
}
func getMacMD5_241019() string {
	id, err := machineid.ID()
	if err != nil {
		return err.Error()
	}
	id = strings.ToLower(id)
	id = strings.ReplaceAll(id, "-", "")
	return id
}

func printAD() {
	ad := Cli.GetAD()
	if len(ad) == 0 {
		return
	}
	fmt.Printf(yellow, ad)
}

func checkUpdate(version int) {
	upUrl := Cli.CheckVersion(fmt.Sprint(version))
	if upUrl == "" {
		return
	}
	isCopyText := ""
	installCmd := `bash -c "$(curl -fsSLk ` + githubPath + `install.sh)"`
	errClip := clipboard.WriteAll(installCmd)
	if errClip == nil {
		isCopyText = Trr.Tr("（已复制到剪贴板）")
	}
	switch runtime.GOOS {
	case "windows":
		fmt.Printf(red, Trr.Tr("有新版本，请关闭本窗口，将下面命令粘贴到GitBash窗口执行")+isCopyText+`：`)
	default:
		fmt.Printf(red, Trr.Tr("有新版本，请关闭本窗口，将下面命令粘贴到新终端窗口执行")+isCopyText+`：`)
	}
	fmt.Printf(hGreen, installCmd)
	_, _ = fmt.Scanln()
	os.Exit(0)
	return
}

// 获取推广人
func getPromotion() (promotion string) {
	b, _ := os.ReadFile(os.Getenv("HOME") + "/.cursor-viprc")
	promotion = strings.TrimSpace(string(b))
	if len(promotion) == 0 {
		if len(os.Args) > 1 {
			promotion = os.Args[1]
		}
	}
	return
}

func getLocale() (langRes, locRes string) {
	osHost := runtime.GOOS
	langRes = "en"
	locRes = "US"
	switch osHost {
	case "windows":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			langRes = langLoc[0]
			langRes = strings.Split(langRes, "-")[0]
			locRes = langLoc[1]
			return
		}
	case "darwin":
		// Exec shell Get-Culture on MacOS.
		cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			langRes = langLoc[0]
			langRes = strings.Split(langRes, "-")[0]
			if len(langLoc) == 1 {
				return
			}
			locRes = langLoc[1]
			return
		}
		plistB, err := os.ReadFile(os.Getenv("HOME") + "/Library/Preferences/.GlobalPreferences.plist")
		if err != nil {
			fmt.Println(err)
			_, _ = fmt.Scanln()
			panic(err)
		}
		var a map[string]interface{}
		_, err = plist.Unmarshal(plistB, &a)
		if err != nil {
			fmt.Println(err)
			_, _ = fmt.Scanln()
			panic(err)
		}
		langLocRaw := a["AppleLocale"].(string)
		langLoc := strings.Split(langLocRaw, "_")
		langRes = langLoc[0]
		langRes = strings.Split(langRes, "-")[0]
		if len(langLoc) == 1 {
			return
		}
		locRes = langLoc[1]
		return
	case "linux":
		envlang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.TrimSpace(envlang)
			langLocRaw = strings.Split(envlang, ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			langRes = langLoc[0]
			langRes = strings.Split(langRes, "-")[0]
			if len(langLoc) == 1 {
				return
			}
			locRes = langLoc[1]
			return
		}
	}
	if langRes == "" {
		langLocRaw := os.Getenv("LC_CTYPE")
		langLocRaw = strings.Split(langLocRaw, ".")[0]
		langLoc := strings.Split(langLocRaw, "_")
		langRes = langLoc[0]
		langRes = strings.Split(langRes, "-")[0]
		if len(langLoc) == 1 {
			return
		}
		locRes = langLoc[1]
		return
	}
	return
}
