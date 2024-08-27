package tui

import (
	"crypto/md5"
	"embed"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"howett.net/plist"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/unknwon/i18n"
)

var version = 100

var hosts = []string{"http://129.154.205.7:7193"}
var host = hosts[0]
var githubPath = "https://mirror.ghproxy.com/https://github.com/kingparks/cursor-vip/releases/download/latest/"
var err error

var green = "\033[32m%s\033[0m\n"
var yellow = "\033[33m%s\033[0m\n"
var hGreen = "\033[1;32m%s\033[0m"
var dGreen = "\033[4;32m%s\033[0m\n"
var red = "\033[31m%s\033[0m\n"
var defaultColor = "%s"
var lang, _ = getLocale()
var deviceID = getMacMD5()
var Cli = Client{Hosts: hosts}

//go:embed all:locales
var localeFS embed.FS

type Tr struct {
	i18n.Locale
}

var Trr *Tr

var jbProduct = []string{"cursor IDE"}

func Run() (productIndexSelected string) {
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
	Cli.SetProxy(lang)
	sCount, sPayCount, _, _, exp := Cli.GetMyInfo(deviceID)
	fmt.Printf(green, Trr.Tr("设备码")+":"+deviceID)
	expTime, _ := time.ParseInLocation("2006-01-02 15:04:05", exp, time.Local)
	fmt.Printf(green, Trr.Tr("付费到期时间")+":"+exp)
	fmt.Printf("\033[32m%s\033[0m\u001B[1;32m %s \u001B[0m\033[32m%s\033[0m\u001B[1;32m %s \u001B[0m\u001B[32m%s\u001B[0m\n",
		Trr.Tr("推广命令：(已推广"), sCount, Trr.Tr("人,推广已付费"), sPayCount, Trr.Tr("人；每推广10人或推广付费2人可获得一年授权)"))
	fmt.Printf(hGreen, "bash <(curl "+githubPath+"install.sh) "+deviceID+"\n")

	printAD()
	checkUpdate(version)
	fmt.Println()
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
			productIndexSelected += jbProduct[vi-1] + ","
		}
		if len(productIndexSelected) > 1 {
			productIndexSelected = productIndexSelected[:len(productIndexSelected)-1]
		}
		fmt.Println(Trr.Tr("选择的产品为：") + productIndexSelected)
		fmt.Println()
	} else {
		productIndexSelected = jbProduct[0]
	}
	// 到期了
	if expTime.Before(time.Now()) {
		payUrl, orderID := Cli.GetPayUrl()
		isCopyText := ""
		errClip := clipboard.WriteAll(payUrl)
		if errClip == nil {
			isCopyText = Trr.Tr("（已复制到剪贴板）")
		}
		fmt.Println(Trr.Tr("使用浏览器打开下面地址进行捐赠") + isCopyText)
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
		isOk, result := Cli.GetLic()
		if !isOk {
			fmt.Printf(red, result)
			return
		}
		fmt.Println()
	}
	fmt.Printf(green, Trr.Tr("授权成功！使用过程请不要关闭此窗口"))
	return
}

func getMacMD5() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	var macAddress []string
	for _, inter := range interfaces {
		// 大于en6的排除
		if strings.HasPrefix(inter.Name, "en") {
			numStr := inter.Name[2:]
			num, _ := strconv.Atoi(numStr)
			if num > 6 {
				continue
			}
		}
		if strings.HasPrefix(inter.Name, "en") || strings.HasPrefix(inter.Name, "Ethernet") || strings.HasPrefix(inter.Name, "以太网") || strings.HasPrefix(inter.Name, "WLAN") {
			macAddress = append(macAddress, inter.HardwareAddr.String())
		}
	}
	sort.Strings(macAddress)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(macAddress, ","))))
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
	if upUrl != "" {
		fmt.Printf(red, Trr.Tr("有新版本可用，尝试自动更新中，若失败，请输入下面命令并回车手动更新程序："))
		fmt.Println()
		fmt.Println(`bash -c "$(curl -fsSL ` + githubPath + `install.sh)"`)
		var cmd *exec.Cmd
		if strings.Contains(strings.ToLower(os.Getenv("ComSpec")), "cmd.exe") {
			cmd = exec.Command("C:\\Program Files\\Git\\git-bash.exe", "-c", fmt.Sprintf(`bash -c "$(curl -fsSL %sinstall.sh)"`, githubPath))
		} else {
			cmd = exec.Command("bash", "-c", fmt.Sprintf(`bash -c "$(curl -fsSL %sinstall.sh)"`, githubPath))
		}
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Trr.Tr("更新完成，重新运行程序即可"))
		os.Exit(0)
		return
	}
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
			panic(err)
		}
		var a map[string]interface{}
		_, err = plist.Unmarshal(plistB, &a)
		if err != nil {
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
